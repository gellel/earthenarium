package latitude

import "fmt"

func NewLatitude(latitude float32) *Latitude {
	ok := ((latitude >= Minimum) && (latitude <= Maximum))
	if ok != true {
		panic(fmt.Errorf(errorLatitude, latitude, Maximum, Minimum))
	}
	l := Latitude(latitude)
	return &l
}
