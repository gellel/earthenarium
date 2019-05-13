package longtitude

import (
	"math/rand"
	"time"
)

func Random() *Longtitude {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewLongtitude(Minimum + (rand.Float32() * ((Maximum - Minimum) + 1)))
}
