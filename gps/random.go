package gps

import (
	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

func Random() *GPS {
	latitude := latitude.Random()
	longtitude := longtitude.Random()
	elevation := elevation.Random()
	hemisphere := hemisphere.NewHemisphere(latitude, longtitude)
	return &GPS{
		Elevation:  elevation,
		Hemisphere: hemisphere,
		Latitude:   latitude,
		Longtitude: longtitude}
}
