package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

type Hemisphere struct {
	X string
	Y string
}

func NewHemisphere(latitude *latitude.Latitude, longitude *longtitude.Longtitude) *Hemisphere {
	return &Hemisphere{
		X: longitude.Hemisphere(),
		Y: latitude.Hemisphere()}
}
