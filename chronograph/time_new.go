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
	local, err := time.Parse(Layout, Default)
	if ok := err == nil; !ok {
		panic(err)
	}
	UTC, _ := time.Parse(Layout, timestamp)
	if ok := err == nil; !ok {
		panic(err)
	}
	/* deconstruct timestamp components */
	nanosecond := UTC.Nanosecond()
	second := UTC.Second()
	minute := UTC.Minute()
	hour := UTC.Hour()
	day := UTC.Day()
	month := UTC.Month()
	year := UTC.Year()
	/* localize dummy timestamp */
	local = local.In(location)
	/* remove required minimum padding from time */
	local = local.Add(time.Duration(-nanosecond) * time.Nanosecond)
	local = local.Add(time.Duration(-second) * time.Second)
	local = local.Add(time.Duration(-minute) * time.Minute)
	local = local.Add(time.Duration(-hour) * time.Hour)
	/* add time from argument timestamp */
	local = local.AddDate(year, int(month)-1, day)
	local = local.Add(time.Duration(nanosecond) * time.Nanosecond)
	local = local.Add(time.Duration(second) * time.Second)
	local = local.Add(time.Duration(minute) * time.Minute)
	local = local.Add(time.Duration(hour) * time.Hour)
	return &Time{
		Day:        NewDay(day, local.YearDay(), local.Weekday()),
		Location:   location,
		Minute:     minute,
		Month:      NewMonth(month),
		Nanosecond: nanosecond,
		Second:     second,
		Time:       &local,
		Unix:       local.Unix(),
		Year:       local.Year()}
}
