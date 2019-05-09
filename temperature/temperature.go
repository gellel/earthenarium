package temperature

import (
	"github.com/gellel/earthenarium/coordinate"
	"github.com/gellel/earthenarium/season"
)

const (
	U = C
)

const (
	C string = "Celcius"
	F string = "Fahrenheit"
	K string = "Kelvin"
)

const (
	Celcius    float32 = 0.0
	Fahrenheit float32 = 32.0
	Kelvin     float32 = 273.15
)

const (
	Average float32 = 14.0
	Maximum float32 = 31.0
	Minimum float32 = -41.0
)

var (
	AboveZero = fibonacci(0, 2.025, 7)
	BelowZero = fibonacci(0, -2.025, 7)
)

func fibonacci(a, b float32, n int) []float32 {
	x := []float32{}
	for i := 0; i < n; i++ {
		T := a
		a = b
		b = T + a
		x = append(x, b)
	}
	return x
}

func NewTemperature(coordinate *coordinate.Coordinate, season *season.Season) *Temperature {

	return &Temperature{}
}

type Temperature struct {
	Max float32
	Min float32
}

func (temperature *Temperature) Celcius() (float32, float32) {
	return temperature.Max, temperature.Min
}

func (temperature *Temperature) Fahrenheit() (float32, float32) {
	return ((temperature.Max * 9 / 5) + Fahrenheit), ((temperature.Min * 9 / 5) + Fahrenheit)
}

func (temperature *Temperature) Kelvin() (float32, float32) {
	return (temperature.Max * Kelvin), (temperature.Min * Kelvin)
}
