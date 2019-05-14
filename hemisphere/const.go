package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

const (
	divison float32 = 0.00001
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
	northLabel string = "northern hemisphere"
	southLabel string = "southern hemisphere"
)

const (
	eastLabel string = "eastern hemisphere"
	westLabel string = "western hemisphere"
)

const (
	equatorLabel string = "equator"
	primeLabel   string = "prime meridian"
)

const (
	antarcticLabel string = "antarctic circle"
	arcticLabel    string = "arctic circle"
)

const (
	cancerLabel    string = "tropic of cancer"
	capricornLabel string = "tropic of capricorn"
)

const (
	N string = northLabel
)
const (
	E string = eastLabel
)
const (
	S string = southLabel
)
const (
	W string = westLabel
)
