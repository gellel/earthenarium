package main

import (
	"fmt"
	"os"

	"github.com/gellel/earthenarium/time"
)

func main() {

	fmt.Println(time.Argparse(os.Args[1], os.Args[2], os.Args[3]))

	/*
		fmt.Println(time.Now().ISOWeek())

		x, _ := time.LoadLocation("Australia/Sydney")
		fmt.Println(x)

		season := season.NewSeasonNorth(chronograph.NewTimeFromISO("2015-01-01T11:11:59.000Z"))

		// coordinate := coordinate.NewCoordinate(44.1, 112.3, (80+180)/2)

		//	temperature := temperature.NewTemperature(coordinate, season)

		fmt.Println(season.Aforetime(), season.Hindmost())

		fmt.Println(temperature.AboveZero, temperature.BelowZero)
	*/
}
