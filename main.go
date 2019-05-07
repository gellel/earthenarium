package main

import (
	"fmt"
	"time"

	"github.com/gellel/earthenarium/chronograph"
)

func main() {

	str := "2015-12-23T05:02:12Z"
	t := time.Now()

	a := chronograph.NewTimeFromISO(str)
	b := chronograph.NewTime(&t)

	d := chronograph.NewSpan(a, b)

	fmt.Println(d.Days)

	fmt.Println(a.Month.Name)
}
