package distance

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
)

// With the inaccuracy coming from the value of the earth's radius
// and the representation of floats, I consider a result valid if it's whithin
// 500 meters from the result
const closeEnough = 0.5

// Data to validate the algorithm
// Verified with http://www.movable-type.co.uk/scripts/latlong.html
var SampleData = map[[4]float64]float64{
	[4]float64{0, 0, 0, 0}:                                      0,
	[4]float64{41.909147, -72.45026, 44.7793732, -63.6734886}:   777.6,
	[4]float64{39.768434, -104.901872, 44.7793732, -63.6734886}: 3400,

	// Add more points here
}

func TestDistanceCalculation(t *testing.T) {
	for latlongs, distance := range SampleData {
		pos1 := &Position{latlongs[0], latlongs[1]}
		pos2 := &Position{latlongs[2], latlongs[3]}

		result := HaversineDistance(pos1, pos2, false)

		if result-distance > closeEnough || distance-result > closeEnough {
			t.Fatalf("Distances not matching. Expecting %f, got %f", distance, result)
		}

	}
}

func BenchmarkHaversineKm(b *testing.B) {

	benchPositions, err := loadCSV()
	if err != nil {
		b.Fatalf(err.Error())
	}
	count := len(benchPositions)
	for i := 0; i < b.N; i++ {
		index := i % count
		HaversineDistance(benchPositions[index][0], benchPositions[index][1], false)
	}

}

func BenchmarkHaversineMiles(b *testing.B) {

	benchPositions, err := loadCSV()
	if err != nil {
		b.Fatalf(err.Error())
	}
	count := len(benchPositions)
	for i := 0; i < b.N; i++ {
		index := i % count
		HaversineDistance(benchPositions[index][0], benchPositions[index][1], true)
	}

}

func loadCSV() ([][2]*Position, error) {
	var positions [][2]*Position
	file, err := os.Open("gis-data.csv")
	if err != nil {
		return positions, fmt.Errorf("couldn't open csv file.")
	}

	csvFile := csv.NewReader(file)
	data, err := csvFile.ReadAll()
	if err != nil || len(data) < 1 {
		return positions, fmt.Errorf("couldn't read csv file. Check the formatting")
	}
	// skip the field denfition row
	data = data[1:]

	for _, row := range data {
		lat1, _ := strconv.ParseFloat(row[1], 64)
		long1, _ := strconv.ParseFloat(row[2], 64)
		lat2, _ := strconv.ParseFloat(row[3], 64)
		long2, _ := strconv.ParseFloat(row[4], 64)

		positions = append(positions, [2]*Position{
			&Position{lat1, long1},
			&Position{lat2, long2},
		})
	}
	return positions, nil
}
