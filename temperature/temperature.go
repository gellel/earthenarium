package temperature

import (
	"fmt"
	"math"
	"time"

	"github.com/gellel/earthenarium/chronograph"
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

func (temperature *Temperature) Time(t *chronograph.Time) *Temperature {
	v := float64(*temperature)
	max := v * 1.2
	min := v * 0.7
	if v < 0 {
		min, max = max, min
	}
	start := chronograph.NewTime(time.Date(t.Year.Number, t.Month.Literal, t.Day.Number, 0, 0, 0, 0, t.Location))
	span := chronograph.NewTimespan(start, t)
	minutes := float64(span.Minutes)
	celcius := math.Abs((max - min) / (12 * 60))
	n := max - ((12*60)-minutes)*(celcius)
	return NewTemperature(float32(n))
}

func (temperature *Temperature) Value() float32 {
	return float32(*temperature)
}
