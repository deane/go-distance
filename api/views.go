//  {
//    "places" :[
//      {
//        "name": <name>,
//        "latitude": <lat>,
//        "longitude": <long>
//      },
//      {
//        "name": <name>,
//        "latitude": <lat>,
//        "longitude": <long>
//      },
//      ...
//    ]
//  }

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/deane/go-distance/distance"
	"github.com/gorilla/mux"
)

// the format expeceted from the JSON requests
type requestFormat struct {
	Places []*distance.Position
}

func calculateDistance(rw http.ResponseWriter, req *http.Request) {
	requestPositions := &requestFormat{}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(jsonError(1, "Houston, we have a problem"))
		return
	}

	// decode the request body
	err = json.Unmarshal(body, requestPositions)
	if err != nil || len(requestPositions.Places) != 2 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(jsonError(2, "Bad Request, wrong format"))
		return
	}

	unit := mux.Vars(req)["unit"]
	useMiles := unit == "miles"

	distance := distance.HaversineDistance(
		requestPositions.Places[0], requestPositions.Places[1], useMiles,
	)
	response, _ := json.Marshal(map[string]interface{}{
		"distance": distance,
		"unit":     unit,
	})
	rw.Write(response)
}
