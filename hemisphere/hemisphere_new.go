package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

func NewHemisphere(latitude *latitude.Latitude, longtitude *longtitude.Longtitude) Hemisphere {
	return Hemisphere{
		Latitude:   NewLatitude(latitude),
		Longtitude: NewLongtitude(longtitude)}
}
