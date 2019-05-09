package coordinate

import (
	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

func NewCoordinate(lat, long, alt float32) *Coordinate {
	return &Coordinate{
		Elevation:  elevation.NewElevation(alt),
		Latitude:   latitude.NewLatitude(lat),
		Longtitude: longtitude.NewLongtitude(long)}
}

type Coordinate struct {
	Elevation  *elevation.Elevation
	Latitude   *latitude.Latitude
	Longtitude *longtitude.Longtitude
}
