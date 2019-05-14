package temperature_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/chronograph"

	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/temperature"
)

func TestNewTemperature(testing *testing.T) {

	elevation := elevation.NewElevation(5000)
	temperature := temperature.NewTemperature(15)
	fmt.Println(temperature.Lapse(elevation))

	t := chronograph.NewTime("2019-01-11T10:01:01.000Z", "australia", "sydney")

	temperature.Diurnal(t)
}
