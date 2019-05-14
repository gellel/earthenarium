package main

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/gps"
	"github.com/gellel/earthenarium/season"
)

func main() {

	t := chronograph.NewTimeLocal("2016-04-01T01:10:00.000Z", "australia", "sydney")

	g := gps.Random()

	fmt.Println(t.Year.Leap)

	fmt.Println(season.NewSeason(t, g.Hemisphere).Ends)
}
