package coordinate

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

type Coordinate struct {
	Latitude   *latitude.Latitude
	Longtitude *longtitude.Longtitude
}

func NewCoordinate(lat, long float32) *Coordinate {
	return &Coordinate{
		Latitude:   latitude.NewLatitude(lat),
		Longtitude: longtitude.NewLongtitude(long)}
}
