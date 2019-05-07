package main

import (
	"fmt"

	"github.com/gellel/earthenarium/longtitude"
)

func main() {
	/*
		time.March

		str := "2015-12-23T05:02:12Z"
		t := time.Now()

		a := chronograph.NewTimeFromISO(str)
		b := chronograph.NewTime(&t)

		d := chronograph.NewSpan(a, b)

		fmt.Println(d.Days)

		fmt.Println(a.Zone)
		x1 := longtitude.NewLongtitude(179.9)

		fmt.Println(x1.ToEast())
		fmt.Println(x1.ToWest())

		x2 := longtitude.NewLongtitude(-1.00)

		fmt.Println(x2.ToEast())
		fmt.Println(x2.ToWest())
	*/

	fmt.Println(longtitude.PrimeMeridian)

	fmt.Println(longtitude.Meridians)

	a := longtitude.NewLongtitude(121.311)
	b := longtitude.NewLongtitude(123.2323)

	c := a.From(b)

	fmt.Println(c.Value())
}
