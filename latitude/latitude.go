package latitude

import (
	"fmt"
	"math"
)

// Latitude stores a numeric coordinate referencing a celestial bodies Y axis.
type Latitude float32

// Absolute returns the numeric value held by the Latitude pointer to an absolute number.
func (latitude *Latitude) Absolute() float32 {
	return float32(math.Abs(float64(*latitude)))
}

// Correct returns a boolean that identifies whether the Latitude value does not exceed the accepted Latitude bounds.
func (latitude *Latitude) Correct() bool {
	return (latitude.Value() >= Minimum) && (latitude.Value() <= Maximum)
}

// Float64 returns a the Latitude value as a 64 bit floating number.
func (latitude *Latitude) Float64() float64 {
	return float64(*latitude)
}

// From returns a Latitude expressing the distance between to Latitude pointers, using the current Latitude as the subtraction.
func (latitude *Latitude) From(l *Latitude) *Latitude {
	return NewLatitude(l.Value() - latitude.Value())
}

// Max returns a new Latitude pointer containing the largest sum of the two Latitudes.
func (latitude *Latitude) Max(l *Latitude) *Latitude {
	return NewLatitude(float32(math.Max(math.Abs(float64(*latitude)), math.Abs(float64(*l)))))
}

// Measurement returns the measurement unit used by the Latitude pointer.
func (latitude *Latitude) Measurement() string {
	return Measurement
}

// Min returns a new Latitude pointer containing the smallest sum of the two Latitudes.
func (latitude *Latitude) Min(l *Latitude) *Latitude {
	return NewLatitude(float32(math.Min(math.Abs(float64(*latitude)), math.Abs(float64(*l)))))
}

func (latitude *Latitude) String() string {
	return fmt.Sprintf("%v", float32(*latitude))
}

// To returns a Latitude expressing the distance between two Latitude pointers, using the argument Latitude as the subtraction.
func (latitude *Latitude) To(l *Latitude) *Latitude {
	return NewLatitude(latitude.Value() - l.Value())
}

// Value returns the numeric value held by the Latitude pointer.
func (latitude *Latitude) Value() float32 {
	return float32(*latitude)
}
