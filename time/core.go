package time

import (
	"time"
	t "time"
)

func Argparse(timestamp, zone, city string) time.Time {
	location, err := t.LoadLocation(NewLocation(zone, city).String())
	ok := err == nil
	if ok != true {
		panic(err)
	}
	time, err := time.Parse(Layout, "0000-01-01T00:00:00.000Z")
	ok = err == nil
	if ok != true {
		panic(err)
	}
	time = time.In(location)
	time = time.Add(t.Duration(-time.Nanosecond()) * t.Nanosecond)
	time = time.Add(t.Duration(-time.Second()) * t.Second)
	time = time.Add(t.Duration(-time.Minute()) * t.Minute)
	time = time.Add(t.Duration(-time.Hour()) * t.Hour)
	UTC, _ := t.Parse(Layout, timestamp)
	time = time.AddDate(UTC.Year(), int(UTC.Month()-1), UTC.Day())
	time = time.Add(t.Duration(UTC.Nanosecond()) * t.Nanosecond)
	time = time.Add(t.Duration(UTC.Second()) * t.Second)
	time = time.Add(t.Duration(UTC.Minute()) * t.Minute)
	time = time.Add(t.Duration(UTC.Hour()) * t.Hour)
	return time
}
