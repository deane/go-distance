package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/deane/go-distance/distance"
)

// callDistance posts the provided string to the /distance/km endpoint
func callDistance(json string) (*http.Response, error) {
	server := httptest.NewServer(newRouter())
	defer server.Close()

	url := fmt.Sprintf("%s/distance/km", server.URL)
	return http.DefaultClient.Post(url, "application/json", strings.NewReader(json))
}

func TestDistancePostKm(t *testing.T) {
	resp, err := callDistance(`{"places":[{"latitude":-34.5, "longitude":14.8}, {"latitude":65.1344, "longitude":-12.4}]}`)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Error when calling distance endpoint. Status code: %d", resp.StatusCode)
	}
}

func TestDistanceNoPlaces(t *testing.T) {

	resp, err := callDistance(`{"places":[]}`)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected response when giving no places. Status code: %d", resp.StatusCode)
	}
}

func TestDistanceOnePlace(t *testing.T) {

	resp, err := callDistance(`{"places":[{"latitude":-34.5, "longitude":14.8}]}`)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected response when giving one place. Status code: %d", resp.StatusCode)
	}
}

func TestDistanceTooManyPlaces(t *testing.T) {

	resp, err := callDistance(`{"places":[{"latitude":-34.5, "longitude":14.8}, {"latitude":65.1344, "longitude":-12.4}, {"latitude":-34.5, "longitude":14.8}, {"latitude":65.1344, "longitude":-12.4}]}`)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected response when giving more than 2 places. Status code: %d", resp.StatusCode)
	}
}

func BenchmarkDistancePost(b *testing.B) {
	server := httptest.NewServer(newRouter())
	defer server.Close()

	url := fmt.Sprintf("%s/distance/km", server.URL)
	requestsJSON, err := getRequestsJSON()
	if err != nil {
		b.Fatalf(err.Error())
	}

	count := len(requestsJSON)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := i % count
		stringReader := strings.NewReader(requestsJSON[index])
		resp, err := http.DefaultClient.Post(url, "application/json", stringReader)
		if err != nil {
			b.Fatalf(err.Error())
		}
		resp.Body.Close()
	}
}

func loadCSV() ([]*requestFormat, error) {
	var requests []*requestFormat
	file, err := os.Open("gis-data.csv")
	if err != nil {
		return requests, fmt.Errorf("couldn't open csv file.")
	}

	csvFile := csv.NewReader(file)
	data, err := csvFile.ReadAll()
	if err != nil || len(data) < 1 {
		return requests, fmt.Errorf("couldn't read csv file. Check the formatting")
	}
	// skip the field denfition row
	data = data[1:]

	for _, row := range data {
		lat1, _ := strconv.ParseFloat(row[1], 64)
		long1, _ := strconv.ParseFloat(row[2], 64)
		lat2, _ := strconv.ParseFloat(row[3], 64)
		long2, _ := strconv.ParseFloat(row[4], 64)

		requests = append(requests, &requestFormat{
			[]*distance.Position{
				&distance.Position{lat1, long1},
				&distance.Position{lat2, long2},
			},
		})
	}
	return requests, nil
}

func getRequestsJSON() ([]string, error) {
	var requestsJSON []string

	requestsData, err := loadCSV()
	if err != nil {
		return requestsJSON, err
	}

	for _, requestData := range requestsData {
		json, err := json.Marshal(requestData)
		if err != nil {
			return requestsJSON, err
		}
		requestsJSON = append(requestsJSON, string(json))
	}

	return requestsJSON, nil

}
