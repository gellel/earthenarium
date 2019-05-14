package chronograph

import (
	"time"
)

// NewTime instantiates a new Time pointer.
// NewTime takes an input of three time components to generate a localized timestamp.
func NewTime(timestamp, zone, city string) *Time {
	t := NewTimestamp(timestamp)
	if ok := t.Correct(); !ok {
		panic(t.Error())
	}
	z := NewZone(zone)
	if ok := z.Correct(); !ok {
		panic(z.Error())
	}
	c := NewCity(city)
	if ok := c.Correct(); !ok {
		panic(c.Error())
	}
	l := NewLocation(z, c)
	if ok := l.Correct(); !ok {
		panic(c.Error())
	}
	location, err := time.LoadLocation(l.String())
	if ok := err == nil; !ok {
		panic(err)
	}
	year, month, day, hour, minute, second, nanosecond := t.Parse()
	local := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, location)
	return &Time{
		Day:        NewDay(local.Day(), local.YearDay(), local.Weekday()),
		Hour:       local.Hour(),
		Location:   location,
		Minute:     local.Minute(),
		Month:      NewMonth(local.Month()),
		Nanosecond: local.Nanosecond(),
		Second:     local.Second(),
		Literal:    &local,
		Unix:       local.Unix(),
		Year:       local.Year()}
}
