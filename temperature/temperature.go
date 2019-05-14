package temperature

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"

	"github.com/gellel/earthenarium/elevation"
)

type Temperature float32

func (temperature *Temperature) Diurnal(time *chronograph.Time) *Temperature {
	fmt.Println(time.Unix / 60)

	return temperature
}

func (temperature *Temperature) Fahrenheit() float32 {
	return ((float32(*temperature) * 9 / 5) + Fahrenheit)
}

func (temperature *Temperature) Kelvin() float32 {
	return (float32(*temperature) * Kelvin)
}

func (temperature *Temperature) Lapse(elevation *elevation.Elevation) *Temperature {
	return temperature.Set(float32(*temperature) - (Decline * (elevation.Value() * 0.001)))
}

func (temperature *Temperature) Set(t float32) *Temperature {
	*temperature = Temperature(t)
	return temperature
}

func (temperature *Temperature) String() string {
	return fmt.Sprintf("%v", float32(*temperature))
}
func (temperature *Temperature) Value() float32 {
	return float32(*temperature)
}
