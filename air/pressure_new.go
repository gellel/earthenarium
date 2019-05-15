package air

import (
	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/temperature"
)

func NewPressure(temperature *temperature.Temperature, elevation *elevation.Elevation) Pressure {
	return Pressure(float32(calculate(0, temperature, elevation)))
}
