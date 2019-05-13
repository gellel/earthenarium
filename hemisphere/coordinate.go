package hemisphere

import "math"

type Coordinate float32

func (coordinate Coordinate) Absolute() float32 {
	return float32(math.Abs(float64(coordinate)))
}

func (coordinate Coordinate) Less(c Coordinate) bool {
	return (coordinate < c)
}

func (coordinate Coordinate) More(c Coordinate) bool {
	return (coordinate > c)
}

func (coordinate Coordinate) Value() float32 {
	return float32(coordinate)
}
