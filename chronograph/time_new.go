package chronograph

import (
	"time"
)

// NewTime instantiates a new Time pointer.
func NewTime(time time.Time) *Time {
	return &Time{
		Day:        NewDay(time.Day(), time.YearDay(), time.Weekday()),
		Hour:       time.Hour(),
		Location:   time.Location(),
		Minute:     time.Minute(),
		Month:      NewMonth(time.Month()),
		Nanosecond: time.Nanosecond(),
		Second:     time.Second(),
		Literal:    time,
		Unix:       time.Unix(),
		Year:       NewYear(time.Year())}
}

// NewTimeLocal takes an input of three time components to generate a localized timestamp.
func NewTimeLocal(timestamp, zone, city string) *Time {
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
	return NewTime(local)
}
