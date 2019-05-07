package latitude

import (
	"fmt"
	"math"
)

// errorLatitudeFloat32 expresses an unsupported float32 argument to NewLatitude.
const errorLatitudeFloat32 string = "position \"%v\" exceeds Latitude boundaries (%v or %v)."

// Maximum is the maximum float32 Latitude pointers support.
const Maximum float32 = 90.0

// Minimum is the minimum float32 Latitude pointers support.
const Minimum float32 = -90.0

// Latitude expresses the Earth-like latitude of a floating point coordinate.
type Latitude float32

// Safe checks that an argument float32 does not exceed the upper or lower bounds of the supported latitudes.
// Although some positions in Geography may indeed exceed these, for the purposes of
// simplicity these conditions have been normalised to be Earth-like.
func Safe(position float32) bool {
	return ((position >= Minimum) && (position <= Maximum))
}

// NewLatitude instantiates a new Latitude pointer.
func NewLatitude(position float32) *Latitude {
	ok := Safe(position)
	if ok != true {
		panic(fmt.Errorf(errorLatitudeFloat32, position, Minimum, Maximum))
	}
	latitude := Latitude(position)
	return &latitude
}

// Float32 converts the Latitude pointer to a float32.
func (latitude *Latitude) Float32() float32 {
	return float32(*latitude)
}

// Float64 converts the Latitude pointer to a float64.
func (latitude *Latitude) Float64() float64 {
	return float64(*latitude)
}

// Int converts the Latitude pointer to an integer.
func (latitude *Latitude) Int() int {
	return int(*latitude)
}

// IsEquator checks whether the Latitude pointer is situated the equator.
func (latitude *Latitude) IsEquator() bool {
	return (latitude.Int() == 0)
}

// IsNorth checks whether the Latitude pointer is situated in the Northern Hemisphere.
func (latitude *Latitude) IsNorth() bool {
	return latitude.Float32() > 0
}

// IsSouth checks whether the Latitude pointer is situated in the Southern Hemisphere.
func (latitude *Latitude) IsSouth() bool {
	return latitude.Float32() < 0
}

// To computes the relative distance (in degrees) from the current Latitude pointer to the argument Latitude pointer.
func (latitude *Latitude) To(l *Latitude) float32 {
	return latitude.Float32() - l.Float32()
}

// ToEquator computes the current distance (in degrees) from the current Latitude pointer to the Equator ((+/-)0.0).
func (latitude *Latitude) ToEquator() float32 {
	return ((Maximum) - float32(math.Abs(latitude.Float64())))
}

// ToNorthPole computes the current distance (in degrees) from the current Latitude pointer to the North Pole (+90.0).
// Sum is generated as an absolute number, indicating the number of subtractions required to reach the
// geographical-opposite.
func (latitude *Latitude) ToNorthPole() float32 {
	return float32(math.Abs(float64(Maximum - float32(*latitude))))
}

// ToSouthPole computes the current distance (in degrees) from the current Latitude pointer to the South Pole (-90.0).
// Sum is generated as an absolute number, indicating the number of subtractions required to reach the
// geographical-opposite.
func (latitude *Latitude) ToSouthPole() float32 {
	return float32(math.Abs(float64(Minimum - float32(*latitude))))
}
