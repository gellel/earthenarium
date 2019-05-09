package main

import (
	"fmt"

	"github.com/gellel/earthenarium/coordinate"
	"github.com/gellel/earthenarium/season"
	"github.com/gellel/earthenarium/temperature"

	"github.com/gellel/earthenarium/chronograph"
)

func main() {

	season := season.NewSeasonNorth(chronograph.NewTimeFromISO("2015-01-01T11:11:59.000Z"))

	coordinate := coordinate.NewCoordinate(30.1, 112.3, (80+180)/2)

	temperature := temperature.NewTemperature(coordinate, season)

	fmt.Println(season.Aforetime(), season.Hindmost())

	fmt.Println(temperature)
}
