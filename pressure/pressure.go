package pressure

import (
	"fmt"

	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/temperature"
)

func NewPressure(t *temperature.Temperature, e *elevation.Elevation) bool {

	x := ((temperature.Decline * e.Value()) / temperature.Seed.Kelvin())

	fmt.Println(x)

	return false
}
