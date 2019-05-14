package main

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/gps"
	"github.com/gellel/earthenarium/season"
)

func main() {

	t := chronograph.NewTime("2019-01-01T01:10:00.000Z", "australia", "sydney")

	g := gps.Random()

	fmt.Println(season.NewSeason(t, g.Hemisphere).Term)
}
