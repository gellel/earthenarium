package elevation

import "fmt"

func NewElevation(elevation float32) *Elevation {
	ok := ((elevation >= Minimum) && (elevation <= Maximum))
	if ok != true {
		panic(fmt.Errorf(errorElevation, elevation, Maximum, Minimum))
	}
	e := Elevation(elevation)
	return &e
}
