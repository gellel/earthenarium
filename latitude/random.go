package latitude

import (
	"math/rand"
)

func Random() *Latitude {
	return NewLatitude(Minimum + (rand.Float32() * ((Maximum - Minimum) + 1)))
}
