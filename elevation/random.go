package elevation

import "math/rand"

func Random() *Elevation {
	return NewElevation(Minimum + (rand.Float32() * ((Maximum - Minimum) + 1)))
}
