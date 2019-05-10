package time

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func NewCity(city string) City {
	regexp, err := regexp.Compile("\\s{2,}")
	ok := err == nil
	if ok != true {
		panic(err)
	}
	namespace := strings.Replace(city, "_", " ", -1)
	substrings := strings.Split(namespace, "/")
	for i := 0; i < len(substrings); i++ {
		substring := regexp.ReplaceAllString(substrings[i], " ")
		substring = strings.TrimSpace(substring)
		subs := strings.Split(substring, " ")
		for j := 0; j < len(subs); j++ {
			if sub := subs[j]; len(sub) > 3 {
				subs[j] = strings.Title(strings.ToLower(sub))
			}
		}
		substrings[i] = strings.Join(subs, "_")
	}
	return City(strings.Join(substrings, "/"))
}

func NewDay(t *time.Time) Day {
	return Day{
		Minutes:  1400,
		Hours:    24,
		Position: int(t.Weekday()),
		Name:     t.Weekday().String(),
		Number:   t.Day(),
		Seconds:  86400}
}

func NewLocation(zone, city string) Location {
	city = NewCity(city).String()
	zone = NewZone(zone).String()
	namespace := fmt.Sprintf("%s/%s", zone, city)
	return Location(namespace)
}

func NewMonth(t *time.Time) Month {
	month := t.Month()
	begins := time.Date(t.Year(), month, 1, 0, 0, 0, 0, t.Location())
	ends := begins.AddDate(0, 1, -1)
	return Month{
		Begins: NewDay(&begins),
		Days:   ends.Day(),
		Ends:   NewDay(&ends),
		Name:   month.String(),
		Number: int(month)}
}

func NewZone(zone string) Zone {
	regexp, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	namespace := regexp.ReplaceAllString(zone, "")
	namespace = strings.Title(strings.ToLower(namespace))
	if _, ok := Timezones[zone]; ok {
		return Zone("")
	}
	return Zone(namespace)
}
