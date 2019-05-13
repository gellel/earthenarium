package gps

import (
	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

type GPS struct {
	Elevation  *elevation.Elevation
	Hemisphere *hemisphere.Hemisphere
	Latitude   *latitude.Latitude
	Longtitude *longtitude.Longtitude
}
