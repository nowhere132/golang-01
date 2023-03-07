package google_map

import (
	"context"
	"fmt"
	"googlemaps.github.io/maps"
)

func GetLocationByZipCode(c *maps.Client, zipcode string) maps.LatLng {
	r := &maps.GeocodingRequest{
		Address: zipcode,
	}

	res, err := c.Geocode(context.Background(), r)
	if err != nil || len(res) == 0 {
		fmt.Errorf("Bug %v\n", err)
	}

	return res[0].Geometry.Location
}
