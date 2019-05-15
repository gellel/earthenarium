package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/gellel/earthenarium/air"
	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/climate"
	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/humidity"
	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
	"github.com/gellel/earthenarium/season"
	"github.com/gellel/earthenarium/temperature"
)

type area struct {
	city       string
	zone       string
	elevation  float32
	latitude   float32
	longtitude float32
}

type mapset map[int]area

func main() {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	titles := "Location\tPosition\tLocal time\tCondition\tTemperature\tPressure\tHumidity\tSeason\tRegion\tHemisphere"
	fmt.Fprintln(w, titles)

	set := mapset{
		1: area{
			zone:       "australia",
			city:       "sydney",
			latitude:   -33.86,
			longtitude: 151.21,
			elevation:  39},
		2: area{
			zone:       "australia",
			city:       "melbourne",
			latitude:   -37.83,
			longtitude: 144.98,
			elevation:  7},
		3: area{
			zone:       "europe",
			city:       "berlin",
			latitude:   52.520008,
			longtitude: 13.404954,
			elevation:  57},
		4: area{
			zone:       "europe",
			city:       "paris",
			latitude:   48.864716,
			longtitude: 2.349014,
			elevation:  36},
		5: area{
			zone:       "america",
			city:       "New york",
			latitude:   40.730610,
			longtitude: -73.935242,
			elevation:  35},
		6: area{
			zone:       "antarctica",
			city:       "vostok",
			latitude:   78.4645,
			longtitude: 106.8339,
			elevation:  3489},
		7: area{
			zone:       "asia",
			city:       "hong kong",
			latitude:   22.3193,
			longtitude: 114.1694,
			elevation:  13.5},
		8: area{
			zone:       "asia",
			city:       "kathmandu",
			latitude:   27.7172,
			longtitude: 85.3240,
			elevation:  1400},
		9: area{
			zone:       "africa",
			city:       "El aaiun",
			latitude:   27.1418,
			longtitude: -13.18797,
			elevation:  68}}

	now := time.Now()

	rand.Seed(now.UnixNano())

	year := rand.Intn(now.Year()-1899) + 1899
	month := rand.Intn(12-1) + 1
	day := rand.Intn(28-1) + 1

	m := fmt.Sprintf("%v", month)
	d := fmt.Sprintf("%v", day)

	if day < 10 {
		d = fmt.Sprintf("0%s", d)
	}
	if month < 10 {
		m = fmt.Sprintf("0%s", m)
	}

	timestamp := fmt.Sprintf("%v-%s-%sT00:10:00.000Z", year, m, d)

	for _, x := range set {

		t := chronograph.NewTimeLocal(timestamp, x.zone, x.city)

		latitudeSeed := latitude.NewLatitude(x.latitude)

		longtitudeSeed := longtitude.NewLongtitude(x.longtitude)

		hemisphereSeed := hemisphere.NewHemisphere(latitudeSeed, longtitudeSeed)

		elevationSeed := elevation.NewElevation(x.elevation)

		seasonSeed := season.NewSeason(t, hemisphereSeed)

		timeSpan := chronograph.NewTimespan(t, seasonSeed.Ends)

		key := hemisphereSeed.Latitude.Area.Key()

		coordinate := hemisphereSeed.Latitude.Area.Lat(latitudeSeed)

		climateSeed := climate.Zones.Fetch(key)

		climateSeasons := climateSeed.Fetch(coordinate)

		seasonalTemperature := climateSeasons.Fetch(seasonSeed.Name)

		temperatureRange := temperature.NewTemperatures(seasonalTemperature, 5.25023, timeSpan.Days)

		temperatureAverage := (temperatureRange.Average() + (*temperatureRange)[t.Day.Number].Value()) / 2

		temperatureSeed := temperature.NewTemperature(temperatureAverage).Elevation(elevationSeed) //.Time(time)

		pressureSeed := air.NewPressure(temperatureSeed, elevationSeed)

		hpa := pressureSeed.Hpa()

		humidity := humidity.NewHumidity(hemisphereSeed)

		condition := "Sunny"

		if temperatureSeed.Value() < -5 && humidity > 45 {
			condition = "Snow"
		} else if humidity > 70 {
			condition = "Rain"
		}

		localtime := fmt.Sprintf("%v-%v-%v %v:%v:%v", t.Year.Number, t.Month.Number, t.Day.Number, t.Hour, t.Minute, t.Second)

		latlongele := fmt.Sprintf("%v,%v,%v", latitudeSeed.Value(), longtitudeSeed.Value(), elevationSeed.Value())

		h := fmt.Sprintf("%.1f", humidity)
		fields := "%s\t%s\t%s\t%s\t%s\t%v\t%s\t%s\t%s\t%s"
		template := fmt.Sprintf(fields, strings.Title(x.city), latlongele, localtime, condition, temperatureSeed.String(), hpa, h, strings.Title(seasonSeed.Name), strings.Title(hemisphereSeed.Latitude.Name), strings.Title(strings.Replace(hemisphereSeed.Latitude.Reference, "hemisphere", "", -1)))

		fmt.Fprintln(w, template)

	}
	w.Flush()
}
