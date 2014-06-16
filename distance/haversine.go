package distance

import "math"

// The mean radius of the earth in km
const EarthRadius = 6371

// 1 km in miles
const kmToMiles = 0.621371192

// data is in degrees
type Position struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

// HaversineDistance returns the distance between the two positions given
// The distance is calculated using the Haversine formula
// For more info, see https://en.wikipedia.org/wiki/Haversine_formula
// The result is returned in Miles if useMiles is true. Kilometers otherwise.
func HaversineDistance(point1 *Position, point2 *Position, useMiles bool) float64 {

	// convert coordinates into radians
	lat1 := degToRad(point1.Lat)
	long1 := degToRad(point1.Long)
	lat2 := degToRad(point2.Lat)
	long2 := degToRad(point2.Long)

	// the haversine formula
	distance := 2 * EarthRadius * math.Asin(math.Sqrt(
		haversine(lat2-lat1)+math.Cos(lat1)*math.Cos(lat2)*haversine(long2-long1),
	))

	if useMiles {
		return distance * kmToMiles
	} else {
		return distance
	}
}

// haversine is a trigonomic helper function
func haversine(theta float64) float64 {
	return (1 - math.Cos(theta)) / 2
}

// degToRad converts an angle from degrees to radians
func degToRad(angle float64) float64 {
	return angle * math.Pi / 180
}
