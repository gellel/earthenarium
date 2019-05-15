package main

import (
	"fmt"

	"github.com/gellel/earthenarium/elevation"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/climate"
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
	"github.com/gellel/earthenarium/season"
	"github.com/gellel/earthenarium/temperature"
)

func main() {

	time := chronograph.NewTimeLocal("2016-01-01T01:10:00.000Z", "europe", "moscow")

	latitudeSeed := latitude.NewLatitude(-33.86)

	longtitudeSeed := longtitude.NewLongtitude(151.21)

	hemisphereSeed := hemisphere.NewHemisphere(latitudeSeed, longtitudeSeed)

	elevationSeed := elevation.NewElevation(89)

	seasonSeed := season.NewSeason(time, hemisphereSeed)

	timeSpan := chronograph.NewTimespan(time, seasonSeed.Ends)

	key := hemisphereSeed.Latitude.Area.Key()

	coordinate := hemisphereSeed.Latitude.Area.Lat(latitudeSeed)

	climateSeed := climate.Zones.Fetch(key)

	climateSeasons := climateSeed.Fetch(coordinate)

	seasonalTemperature := climateSeasons.Fetch(seasonSeed.Name)

	temperatureRange := temperature.NewTemperatures(seasonalTemperature, 5.25023, timeSpan.Days)

	temperatureAverage := (temperatureRange.Average() + (*temperatureRange)[time.Day.Number].Value()) / 2

	temperatureSeed := temperature.NewTemperature(temperatureAverage)

	fmt.Println(temperatureSeed.Elevation(elevationSeed))
}
