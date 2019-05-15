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
	NorthLabel string = "northern hemisphere"
	SouthLabel string = "southern hemisphere"
)

const (
	EastLabel string = "eastern hemisphere"
	WestLabel string = "western hemisphere"
)

const (
	EquatorLabel string = "equator"
	PrimeLabel   string = "prime meridian"
)

const (
	AntarcticLabel string = "antarctic circle"
	ArcticLabel    string = "arctic circle"
)

const (
	CancerLabel    string = "tropic of cancer"
	CapricornLabel string = "tropic of capricorn"
)

const (
	N string = NorthLabel
)
const (
	E string = EastLabel
)
const (
	S string = SouthLabel
)
const (
	W string = WestLabel
)
