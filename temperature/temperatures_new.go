package temperature

import "math/rand"

func NewTemperatures(mu, sigma float32, x int) *Temperatures {
	t := Temperatures{}
	for i := 0; i < x; i++ {
		n := float32(rand.NormFloat64()*float64(sigma) + float64(mu))
		t = append(t, NewTemperature(n))
	}
	return &t
}
