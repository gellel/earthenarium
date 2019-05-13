package chronograph

import "fmt"

// NewLocation instantiates a new Location pointer.
func NewLocation(zone *Zone, city *City) *Location {
	substring := fmt.Sprintf("%s/%s", zone, city)
	location := Location(substring)
	return &location
}
