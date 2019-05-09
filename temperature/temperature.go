package temperature

import (
	"github.com/gellel/earthenarium/coordinate"
	"github.com/gellel/earthenarium/season"
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
	Maximum float32 = 44.0
	Minimum float32 = -44.0
)

func NewTemperature(coordinate *coordinate.Coordinate, season *season.Season) *Temperature {
	coordinate.Latitude.Float64()

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
