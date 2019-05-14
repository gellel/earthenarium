package temperature

import (
	"fmt"

	"github.com/gellel/earthenarium/elevation"
)

type Temperature float32

func (temperature *Temperature) Elevation(elevation *elevation.Elevation) *Temperature {
	return NewTemperature(temperature.Value() - (Decline * elevation.Kilometers().Value()))
}

func (temperature *Temperature) Fahrenheit() float32 {
	return (((temperature.Value() * 9) / 5) + Fahrenheit)
}

func (temperature *Temperature) Kelvin() float32 {
	return (temperature.Value() + Kelvin)
}

func (temperature *Temperature) Set(t float32) *Temperature {
	*temperature = *NewTemperature(t)
	return temperature
}

func (temperature *Temperature) String() string {
	return fmt.Sprintf("%v", temperature.Value())
}

func (temperature *Temperature) Value() float32 {
	return float32(*temperature)
}
