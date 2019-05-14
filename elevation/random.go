package elevation

import (
	"math/rand"
	"time"
)

func Random() *Elevation {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewElevation(Minimum + (rand.Float32() * ((Maximum - Minimum) + 1)))
}
