/*
go-distance is a web service that returns the distance between two coordinates
using the haversine formula.

To install

	go get github.com/deane/go-distance
	go build github.com/deane/go-distance

To run

	cd $GOPATH/src/github.com/deane/go-distance
	go run main.go

To test

	go test ./...

To benchmark

	go test ./... -bench=.

API usage

go-distance exposes a JSON API:
To get a distance between two places, do a POST to

	/distance/<unit>

where "unit" is "km" or "miles". the body of the request, should be a valid JSON
in the following format

  {
	"places": [
	  {
		"latitude": <lat>
		"longitude": <long>
	  },
	  {
		"latitude": <lat>
		"longitude": <long>
	  }
	]
  }

Exactly 2 places are required
The latitude and longitude should be floats.

The response retuned in the following JSON format:

	{
	  "distance": <distance>,
	  "unit": <unit>
	}

In the case of an error, the response body will be:

	{
	  "error_code": <errorCode>,
	  "message": <message>
	}
*/
package main
