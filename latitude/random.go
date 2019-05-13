package latitude

import (
	"math/rand"
	"time"
)

func Random() *Latitude {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewLatitude(Minimum + (rand.Float32() * ((Maximum - Minimum) + 1)))
}
