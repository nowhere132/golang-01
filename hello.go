package main

import (
	"fmt"
	googleMap "golang-01/google-map"
	"googlemaps.github.io/maps"
)

func main() {
	testDistance()
}

func testDistance() {
	arr := make([]maps.LatLng, 10)

	arr[0] = maps.LatLng{
		Lat: 61.2189216614,
		Lng: -149.8503723145,
	}
	arr[1] = maps.LatLng{
		Lat: 61.2190361023,
		Lng: -149.8662872314,
	}

	c := googleMap.CreateMapsClient()

	fmt.Println(googleMap.GetDistanceUsingHaversineFormula(arr[0], arr[1])) // 0.5331270584324534
	fmt.Println(googleMap.GetDistanceUsingGoogleMap(c, arr[0], arr[1]))
}
