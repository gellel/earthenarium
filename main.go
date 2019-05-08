package main

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/season"
)

func main() {

	y := "2015-02-28T23:59:59.000Z"

	c := chronograph.NewTimeFromISO(y)

	timespan := season.NewTimespan(c)

	fmt.Println(timespan.Spans)
}
