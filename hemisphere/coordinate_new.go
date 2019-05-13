package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

func NewCoordinate(coordinate float32) Coordinate {
	return Coordinate(coordinate)
}

func NewCoordinateFromLatitude(latitude *latitude.Latitude) Coordinate {
	return NewCoordinate(latitude.Value())
}

func NewCoordinateFromLongtitude(longtitude *longtitude.Longtitude) Coordinate {
	return NewCoordinate(longtitude.Value())
}
