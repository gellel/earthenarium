package longtitude

import (
	"fmt"
	"math"

	"github.com/gellel/earthenarium/hemisphere"
)

const (
	errorLongtitude string = "longtitude \"%v\" exceeds supported range; cannot eclipse \"%v\" or \"%v\""
)

const (
	// Maximum easternly distance permitted across the celestial bodies X axis.
	Maximum float32 = 180
	// Minimum westernly distance permitted across the celestial bodies X axis.
	Minimum float32 = -180
)

var (
	// PrimeMeridian is the center-most point across the celestial bodies X axis.
	PrimeMeridian = (Maximum + Minimum)
)

var (
	// Meridian is the base degree of distance across the celestial bodies X axis.
	Meridian = 1
	// Meridians is the sum of all Meridians across the celestial body.
	Meridians = ((Maximum - Minimum) * float32(Meridian))
)

// Exists returns a boolean that identifies whether the positional float32 is supported longtitude.
func Exists(longtitude float32) bool {
	return ((longtitude >= Minimum) && (longtitude <= Maximum))
}

// NewLongtitude creates a new Longtitude pointer from the argument float32.
func NewLongtitude(longtitude float32) *Longtitude {
	ok := Exists(longtitude)
	if ok != true {
		panic(fmt.Errorf(errorLongtitude, longtitude, Maximum, Minimum))
	}
	l := Longtitude(longtitude)
	return &l
}

// Longtitude stores a numeric coordinate referencing a celestial bodies X axis.
type Longtitude float32

// Absolute returns the numeric value held by the Longtitude pointer to an absolute number.
func (longtitude *Longtitude) Absolute() float32 {
	return float32(math.Abs(float64(*longtitude)))
}

// Correct returns a boolean that identifiers whether the Longtitude value does not exceed the accepted Longtitude bounds.
func (longtitude *Longtitude) Correct() bool {
	return (longtitude.Value() >= Minimum) && (longtitude.Value() <= Maximum)
}

// East returns a boolean that identifies whether the Longtitude value is situated east of the prime-meridian.
func (longtitude *Longtitude) East() bool {
	return (longtitude.Value() > 0)
}

// From returns a Longtitude expressing the distance between to Longtitude pointers, using the current Latitude as the subtraction.
func (longtitude *Longtitude) From(l *Longtitude) *Longtitude {
	return NewLongtitude(l.Value() - longtitude.Value())
}

// Hemisphere returns a string describing the celestial hemisphere the Longtitute value is situated.
func (longtitude *Longtitude) Hemisphere() string {
	if longtitude.East() {
		return hemisphere.Eastern
	}
	if longtitude.West() {
		return hemisphere.Western
	}
	return hemisphere.Meridian
}

// Meridian returns a boolean that identifies whether the Longtitude value is situated approximately near or at the prime-meridian.
func (longtitude *Longtitude) Meridian() bool {
	return (longtitude.Absolute() < 10)
}

// To returns a Longtitude expressing the distance between two Longtitude pointers, using the argument Latitude as the subtraction.
func (longtitude *Longtitude) To(l *Longtitude) *Longtitude {
	return NewLongtitude(longtitude.Value() - l.Value())
}

// Value returns the numeric value held by the Longtitude pointer.
func (longtitude *Longtitude) Value() float32 {
	return float32(*longtitude)
}

// West returns a boolean that identifies whether the Longitude value is situated west of the prime-meridian.
func (longtitude *Longtitude) West() bool {
	return (longtitude.Value() < 0)
}
