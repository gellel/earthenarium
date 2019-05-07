package latitude

import (
	"fmt"
	"math"
)

const (
	errorLongtitude string = "latitude \"%v\" exceeds supported range; cannot eclipse \"%v\" or \"%v\""
)

const (
	// Maximum easternly distance permitted across the celestial bodies Y axis.
	Maximum float32 = 90
	// Minimum westernly distance permitted across the celestial bodies Y axis.
	Minimum float32 = -90
)

var (
	// Equator is the center-most point across the celestial bodies Y axis.
	Equator = (Maximum + Minimum)
)

var (
	// Meridian is the base degree of distance across the celestial bodies Y axis.
	Meridian = 1
	// Meridians is the sum of all Meridians across the celestial body.
	Meridians = ((Maximum - Minimum) * float32(Meridian))
)

// Exists returns a boolean that identifies whether the positional float32 is supported latitude.
func Exists(latitude float32) bool {
	return ((latitude >= Minimum) && (latitude <= Maximum))
}

// NewLongtitude creates a new Latitude pointer from the argument float32.
func NewLongtitude(latitude float32) *Latitude {
	ok := Exists(latitude)
	if ok != true {
		panic(fmt.Errorf(errorLongtitude, latitude, Maximum, Minimum))
	}
	l := Latitude(latitude)
	return &l
}

// Latitude stores a numeric coordinate referencing a celestial bodies Y axis.
type Latitude float32

// Absolute returns the numeric value held by the Latitude pointer to an absolute number.
func (latitude *Latitude) Absolute() float32 {
	return float32(math.Abs(float64(*latitude)))
}

// Correct returns a boolean that identifiers whether the Latitude value does not exceed the accepted Latitude bounds.
func (latitude *Latitude) Correct() bool {
	return (latitude.Value() >= Minimum) && (latitude.Value() <= Maximum)
}

// Equator returns a boolean that identifies whether the Latitude value is situated approximately near or at the equator.
func (latitude *Latitude) Equator() bool {
	return (latitude.Absolute() == 0)
}

// From returns a Latitude expressing the distance between to Latitude pointers, using the current Latitude as the subtraction.
func (latitude *Latitude) From(l *Latitude) *Latitude {
	return NewLongtitude(l.Value() - latitude.Value())
}

// Hemisphere returns a string describing the celestial hemisphere the Longtitute value is situated.
func (latitude *Latitude) Hemisphere() string {
	north := latitude.North()
	if north == true {
		return "Northern"
	}
	south := latitude.South()
	if south == true {
		return "Southern"
	}
	return "Equator"
}

// North returns a boolean that identifies whether the Latitude value is situated north of the equator.
func (latitude *Latitude) North() bool {
	return (latitude.Value() > 0)
}

// South returns a boolean that identifies whether the Longitude value is situated south of the equator.
func (latitude *Latitude) South() bool {
	return (latitude.Value() < 0)
}

// To returns a Latitude expressing the distance between two Latitude pointers, using the argument Latitude as the subtraction.
func (latitude *Latitude) To(l *Latitude) *Latitude {
	return NewLongtitude(latitude.Value() - l.Value())
}

// Value returns the numeric value held by the Latitude pointer.
func (latitude *Latitude) Value() float32 {
	return float32(*latitude)
}
