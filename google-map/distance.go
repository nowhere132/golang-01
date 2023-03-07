package google_map

import (
	"context"
	"fmt"
	"googlemaps.github.io/maps"
	"math"
)

func GetDistanceUsingGoogleMap(c *maps.Client, p1 maps.LatLng, p2 maps.LatLng) int {
	toString := func(p *maps.LatLng) string {
		return fmt.Sprintf("%f,%f", p.Lat, p.Lng)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      []string{toString(&p1)},
		Destinations: []string{toString(&p2)},
	}

	res, err := c.DistanceMatrix(context.Background(), r)
	if err != nil || len(res.Rows) == 0 {
		fmt.Errorf("Bug %v\n", err)
	}

	return res.Rows[0].Elements[0].Distance.Meters
}

func GetDistanceUsingHaversineFormula(p1 maps.LatLng, p2 maps.LatLng) float64 {
	// ref: https://stackoverflow.com/questions/1502590/calculate-distance-between-two-points-in-google-maps-v3
	const RadiusInMeters = 6378137
	//const MileInMeters = 1600
	const KilometerInMeters = 1000
	var radiusInMiles float64 = RadiusInMeters / KilometerInMeters

	rad := func(x float64) float64 {
		return x * math.Pi / 180
	}
	sqr := func(x float64) float64 {
		return x * x
	}

	// core logic sits here
	dLat, dLong := rad(p2.Lat-p1.Lat), rad(p2.Lng-p1.Lng)
	a := sqr(math.Sin(dLat/2)) + math.Cos(rad(p1.Lat))*math.Cos(rad(p2.Lat))*sqr(math.Sin(dLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radiusInMiles * c
}
