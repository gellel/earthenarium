package longtitude

import (
	"fmt"
	"math"
)

// Longtitude stores a numeric coordinate referencing a celestial bodies X axis.
type Longtitude float32

// Absolute returns the numeric value held by the Longtitude pointer to an absolute number.
func (longtitude *Longtitude) Absolute() float32 {
	return float32(math.Abs(float64(*longtitude)))
}

// Correct returns a boolean that identifies whether the Longtitude value does not exceed the accepted Longtitude bounds.
func (longtitude *Longtitude) Correct() bool {
	return (longtitude.Value() >= Minimum) && (longtitude.Value() <= Maximum)
}

// Float64 returns a the Longtitude value as a 64 bit floating number.
func (longtitude *Longtitude) Float64() float64 {
	return float64(*longtitude)
}

// From returns a Longtitude expressing the distance between to Longtitude pointers, using the current Longtitude as the subtraction.
func (longtitude *Longtitude) From(l *Longtitude) *Longtitude {
	return NewLongtitude(l.Value() - longtitude.Value())
}

// Max returns a new Longtitude pointer containing the largest sum of the two Longtitudes.
func (longtitude *Longtitude) Max(l *Longtitude) *Longtitude {
	return NewLongtitude(float32(math.Max(math.Abs(float64(*longtitude)), math.Abs(float64(*l)))))
}

// Measurement returns the measurement unit used by the Longtitude pointer.
func (longtitude *Longtitude) Measurement() string {
	return Measurement
}

// Min returns a new Longtitude pointer containing the smallest sum of the two Longtitudes.
func (longtitude *Longtitude) Min(l *Longtitude) *Longtitude {
	return NewLongtitude(float32(math.Min(math.Abs(float64(*longtitude)), math.Abs(float64(*l)))))
}

func (longtitude *Longtitude) String() string {
	return fmt.Sprintf("%v", float32(*longtitude))
}

// To returns a Longtitude expressing the distance between two Longtitude pointers, using the argument Longtitude as the subtraction.
func (longtitude *Longtitude) To(l *Longtitude) *Longtitude {
	return NewLongtitude(longtitude.Value() - l.Value())
}

// Value returns the numeric value held by the Longtitude pointer.
func (longtitude *Longtitude) Value() float32 {
	return float32(*longtitude)
}
