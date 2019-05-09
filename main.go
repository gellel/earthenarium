package main

import (
	"fmt"

	"github.com/gellel/earthenarium/season"

	"github.com/gellel/earthenarium/chronograph"
)

func main() {

	y := "2015-02-28T23:59:59.000Z"

	season := season.NewSeasonNorth(chronograph.NewTimeFromISO(y))

	fmt.Println(season)

	fmt.Println(season.Next())

}
