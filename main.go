package main

import (
	"fmt"

	"github.com/gellel/earthenarium/longtitude"

	"github.com/gellel/earthenarium/latitude"

	"github.com/gellel/earthenarium/hemisphere"
)

func main() {
	/*
		zone := "australia"
		city := "sydney"
		a := "2019-01-01T01:00:00.000Z"
		b := "2019-05-01T01:00:00.000Z"

		x := chronograph.NewTime(a, zone, city)
		y := chronograph.NewTime(b, zone, city)
		z := chronograph.NewTimespan(x, y)

		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)

		fmt.Println(hemisphere.Equator)
		fmt.Println(elevation.Measurement)
		fmt.Println(hemisphere.Cancer.North())
	*/

	fmt.Println(hemisphere.NewHemisphere(latitude.Random(), longtitude.Random()))
}
