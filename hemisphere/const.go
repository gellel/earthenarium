package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

const (
	north float32 = latitude.Maximum
	south float32 = latitude.Minimum
)

const (
	east float32 = longtitude.Maximum
	west float32 = longtitude.Minimum
)

const (
	equator float32 = latitude.Mean
	prime   float32 = longtitude.Mean
)

const (
	antarctic float32 = -arctic
	arctic    float32 = 66.56083
)

const (
	cancer    float32 = 23.5
	capricorn float32 = -cancer
)

const (
	northLabel = "northern hemisphere"
	southLabel = "southern hemisphere"
)

const (
	eastLabel = "eastern hemisphere"
	westLabel = "western hemisphere"
)

const (
	equatorLabel = "equator"
	primeLabel   = "prime meridian"
)

const (
	antarcticLabel = "antarctic circle"
	arcticLabel    = "arctic circle"
)

const (
	cancerLabel    = "tropic of cancer"
	capricornLabel = "tropic of capricorn"
)

const (
	N string = northLabel
	E string = eastLabel
	S string = southLabel
	W string = westLabel
)
