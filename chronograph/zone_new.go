package chronograph

import (
	"regexp"
	"strings"
)

// NewZone instantiates a new Zone pointer.
func NewZone(timezone string) *Zone {
	reg := regexp.MustCompile("[^a-zA-z]+")
	timezone = reg.ReplaceAllString(timezone, "")
	timezone = strings.ToLower(timezone)
	timezone = strings.Title(timezone)
	zone := Zone(timezone)
	return &zone
}
