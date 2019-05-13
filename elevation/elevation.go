package elevation

import (
	"fmt"
	"math"
)

// Elevation stores a numeric coordinate referencing a celestial bodies Z axis.
type Elevation float32

// Absolute returns the numeric value held by the Elevation pointer to an absolute number.
func (elevation *Elevation) Absolute() float32 {
	return float32(math.Abs(float64(*elevation)))
}

// Correct returns a boolean that identifies whether the Elevation value does not exceed the accepted Elevation bounds.
func (elevation *Elevation) Correct() bool {
	return (elevation.Value() >= Minimum) && (elevation.Value() <= Maximum)
}

// Float64 returns a the Elevation value as a 64 bit floating number.
func (elevation *Elevation) Float64() float64 {
	return float64(*elevation)
}

// From returns a Elevation expressing the distance between to Elevation pointers, using the current Elevation as the subtraction.
func (elevation *Elevation) From(l *Elevation) *Elevation {
	return NewElevation(l.Value() - elevation.Value())
}

// Max returns a new Elevation pointer containing the largest sum of the two Elevations.
func (elevation *Elevation) Max(l *Elevation) *Elevation {
	return NewElevation(float32(math.Max(math.Abs(float64(*elevation)), math.Abs(float64(*l)))))
}

// Measurement returns the measurement unit used by the Elevation pointer.
func (elevation *Elevation) Measurement() string {
	return Measurement
}

// Min returns a new Elevation pointer containing the smallest sum of the two Elevations.
func (elevation *Elevation) Min(l *Elevation) *Elevation {
	return NewElevation(float32(math.Min(math.Abs(float64(*elevation)), math.Abs(float64(*l)))))
}

func (elevation *Elevation) String() string {
	return fmt.Sprintf("%v", float32(*elevation))
}

// To returns a Elevation expressing the distance between two Elevation pointers, using the argument Elevation as the subtraction.
func (elevation *Elevation) To(l *Elevation) *Elevation {
	return NewElevation(elevation.Value() - l.Value())
}

// Value returns the numeric value held by the Elevation pointer.
func (elevation *Elevation) Value() float32 {
	return float32(*elevation)
}
