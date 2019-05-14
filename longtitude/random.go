package longtitude

import (
	"math/rand"
)

func Random() *Longtitude {
	return NewLongtitude(Minimum + (rand.Float32() * ((Maximum - Minimum) + 1)))
}
