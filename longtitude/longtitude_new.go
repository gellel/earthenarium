package longtitude

import "fmt"

func NewLongtitude(longtitude float32) *Longtitude {
	ok := ((longtitude >= Minimum) && (longtitude <= Maximum))
	if ok != true {
		panic(fmt.Errorf(errorLongtitude, longtitude, Maximum, Minimum))
	}
	l := Longtitude(longtitude)
	return &l
}
