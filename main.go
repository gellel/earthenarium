package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gellel/earthenarium/air"
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

	city := "sydney"

	time := chronograph.NewTimeLocal("2016-11-01T01:10:00.000Z", "australia", city)

	latitudeSeed := latitude.NewLatitude(-33.86)

	longtitudeSeed := longtitude.NewLongtitude(151.21)

	hemisphereSeed := hemisphere.NewHemisphere(latitudeSeed, longtitudeSeed)

	elevationSeed := elevation.NewElevation(39)

	seasonSeed := season.NewSeason(time, hemisphereSeed)

	timeSpan := chronograph.NewTimespan(time, seasonSeed.Ends)

	key := hemisphereSeed.Latitude.Area.Key()

	coordinate := hemisphereSeed.Latitude.Area.Lat(latitudeSeed)

	climateSeed := climate.Zones.Fetch(key)

	climateSeasons := climateSeed.Fetch(coordinate)

	seasonalTemperature := climateSeasons.Fetch(seasonSeed.Name)

	temperatureRange := temperature.NewTemperatures(seasonalTemperature, 5.25023, timeSpan.Days)

	temperatureAverage := (temperatureRange.Average() + (*temperatureRange)[time.Day.Number].Value()) / 2

	temperatureSeed := temperature.NewTemperature(temperatureAverage).Elevation(elevationSeed) //.Time(time)

	pressureSeed := air.NewPressure(temperatureSeed, elevationSeed)

	hpa := pressureSeed.Hpa()

	var max, min float32

	switch hemisphereSeed.Latitude.Name {
	case hemisphere.CapricornLabel, hemisphere.CancerLabel:
		max, min = 40, 70
	case hemisphere.EquatorLabel:
		max, min = 40, 89
	case hemisphere.AntarcticLabel, hemisphere.ArcticLabel:
		max, min = 20, 60
	}

	humidty := min + rand.Float32()*(max-min)

	condition := "Sunny"

	if humidty > 60 {
		if temperatureSeed.Value() > 0 {
			condition = "Rain"
		} else {
			condition = "Snow"
		}
	}

	localtime := fmt.Sprintf("%v-%v-%v %v:%v:%v", time.Year.Number, time.Month.Number, time.Day.Number, time.Hour, time.Minute, time.Second)

	latlongele := fmt.Sprintf("%v,%v,%v", latitudeSeed.Value(), longtitudeSeed.Value(), elevationSeed.Value())

	h := fmt.Sprintf("%.1f", humidty)

	template := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%v\t%s", strings.Title(city), latlongele, localtime, condition, temperatureSeed.String(), hpa, h)

	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintln(w, "Location\tPosition\tLocal time\tCondition\tTemperature\tPressure\tHumidity")
	fmt.Fprintln(w, template)
	fmt.Fprintln(w)
	w.Flush()

}
