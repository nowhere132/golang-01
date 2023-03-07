package google_map

import (
	"googlemaps.github.io/maps"
	"log"
)

func CreateMapsClient() *maps.Client {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCm4dKU1CA6ezOeoCGUuoj833RoRglEtZw"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	return c
}
