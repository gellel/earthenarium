package time

import (
	"fmt"
	"time"
)

// Argparse takes an input of three time components to generate a localized timestamp.
// Timestamp is intended to be received as YYYY-MM-DDTHH-MM-SS.NNNZ.
// Zone is expected to resemble a region of the Earth's geography, supported list is exported as Time.Timezones.
// City is a Title-case underscore-separated name, for example New_York, Isle_of_Man, etc.
func Argparse(timestamp, zone, city string) time.Time {
	l := NewLocation(zone, city)
	ok := l.Correct()
	if ok != true {
		panic(fmt.Errorf("location string is not formatted correctly \"%s\"", l.String()))
	}
	location, err := time.LoadLocation(l.String())
	ok = err == nil
	if ok != true {
		panic(err)
	}
	local, err := time.Parse(Layout, "0000-01-01T00:00:00.000Z")
	ok = err == nil
	if ok != true {
		panic(err)
	}
	UTC, _ := time.Parse(Layout, timestamp)
	ok = err == nil
	if ok != true {
		panic(err)
	}
	/* localize dummy timestamp */
	local = local.In(location)
	/* remove required minimum padding from time */
	local = local.Add(time.Duration(-local.Nanosecond()) * time.Nanosecond)
	local = local.Add(time.Duration(-local.Second()) * time.Second)
	local = local.Add(time.Duration(-local.Minute()) * time.Minute)
	local = local.Add(time.Duration(-local.Hour()) * time.Hour)
	/* add time from argument timestamp */
	local = local.AddDate(UTC.Year(), int(UTC.Month()-1), UTC.Day())
	local = local.Add(time.Duration(UTC.Nanosecond()) * time.Nanosecond)
	local = local.Add(time.Duration(UTC.Second()) * time.Second)
	local = local.Add(time.Duration(UTC.Minute()) * time.Minute)
	local = local.Add(time.Duration(UTC.Hour()) * time.Hour)
	return local
}
