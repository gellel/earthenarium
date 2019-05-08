package main

import (
	"fmt"

	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/season"

	"github.com/gellel/earthenarium/chronograph"
)

func main() {

	y := "2015-02-28T23:59:59.000Z"

	fmt.Println(season.NewSeason(hemisphere.Northern, chronograph.NewTimeFromISO(y)))

}
