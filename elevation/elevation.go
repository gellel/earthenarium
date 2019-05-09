package elevation

import "fmt"

const (
	errorElevation string = "elevation \"%v\" exceeds supported range; cannot eclipse \"%v\" or \"%v\""
)

const (
	// Average elevation (in meters).
	Average float32 = 840
	// Maximum elevation (in meters) permitted on the celestial body Z axis. Number represens highest point on Earth (Mt. Everest).
	Maximum float32 = 8848
	// Minimum elevation (in meters) permitted on the celestial body Z axis. Number represents sea-floor on Earth.
	Minimum float32 = 0
)

// Exists returns a boolean that identifies whether the positional float32 is supported elevation.
func Exists(elevation float32) bool {
	return ((elevation >= Minimum) && (elevation <= Maximum))
}

func NewElevation(elevation float32) *Elevation {
	ok := Exists(elevation)
	if ok != true {
		panic(fmt.Errorf(errorElevation, elevation, Maximum, Minimum))
	}
	e := Elevation(elevation)
	return &e
}

func NewElevationFromMaxMin(max, min float32) *Elevation {
	return NewElevation((max + min) / 2)
}

type Elevation float32

// Value returns the numeric value held by the Evelvation pointer.
func (elevation *Elevation) Value() float32 {
	return float32(*elevation)
}
