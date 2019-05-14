package main

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
	"github.com/gellel/earthenarium/season"
	"github.com/gellel/earthenarium/temperature"
)

func main() {

	time := chronograph.NewTimeLocal("2016-04-01T01:10:00.000Z", "europe", "moscow")

	latitude := latitude.NewLatitude(-33.86)

	longtitude := longtitude.NewLongtitude(151.21)

	hemisphere := hemisphere.NewHemisphere(latitude, longtitude)

	season := season.NewSeason(time, hemisphere)

	s := chronograph.NewTimespan(time, season.Ends)

	for _, i := range *temperature.NewTemperatures(25, 0.25, s.Days) {
		fmt.Println(i)
	}

	fmt.Println(hemisphere.Latitude.Area.Lat(latitude))

}
