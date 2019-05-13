package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

const (
	Meridian Longtitude = Longtitude(longtitude.Maximum - longtitude.Minimum)
	Eastern  Longtitude = Longtitude(longtitude.Minimum)
	Western  Longtitude = Longtitude(longtitude.Maximum)
)

const (
	Antarctic Latitude = Latitude(-66.56083)
	Arctic    Latitude = Latitude(66.56083)
	Cancer    Latitude = Latitude(23.5)
	Capricorn Latitude = Latitude(-23.5)
	Equator   Latitude = Latitude(latitude.Maximum + latitude.Minimum)
	Northern  Latitude = Latitude(latitude.Maximum)
	Southern  Latitude = Latitude(latitude.Minimum)
)

const (
	AntarcticCircle    Label = Label("Antarctic Circle")
	ArcticCircle       Label = Label("Arctic Circle")
	TropicOfCancer     Label = Label("Tropic of Cancer")
	TropicOfCapricorn  Label = Label("Tropic of Capricorn")
	EquatorLine        Label = Label("Equator")
	EasternHemisphere  Label = Label("Eastern Hemisphere")
	NorthernHemisphere Label = Label("Northern Hemisphere")
	SouthernHemisphere Label = Label("Southern Hemisphere")
	WesternHemisphere  Label = Label("Western Hemisphere")
)
