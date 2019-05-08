package main

import (
	"fmt"

	"github.com/gellel/earthenarium/season"
)

func main() {

	/*str := "2015-12-23T05:02:12Z"
	t := time.Now()

	a := chronograph.NewTimeFromISO(str)
	b := chronograph.NewTime(&t)
	s := chronograph.NewSpan(a, b)

	latlong := coordinate.NewCoordinate(81.13, 139.11)

	x := season.NewSeason(latlong, a, b, s)
	//fmt.Println(a.Day.Number, a.Month.Number, a.Year)

	fmt.Println(x)

	fmt.Print(season.GetSeason(latlong, a))*/

	//fmt.Println(chronograph.NewTimeFromISO("0001-01-01T00:00:00Z").Second)
	//now := time.Now()

	//t := chronograph.NewTime(&now)

	fmt.Println(season.Epochs)

	//fmt.Println(season.GetSeason(coordinate.NewCoordinate(89, 135), chronograph.NewTime(&t)))

	//fmt.Println(season.IsSouthernWinter(chronograph.NewTime(&t)))
}
