package temperature

import (
	"math/rand"
	"time"
)

func Random(min, max float32) *Temperature {
	rand.Seed(time.Now().Unix())
	r := min + (rand.Float32() * ((max - min) + 1))
	return NewTemperature(r)
}
