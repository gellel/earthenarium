package time

import "time"

type Time struct {
	City     City
	Day      Day
	Days     int
	Location time.Location
	Month    Month
	Time     time.Time
	Unix     int64
	Year     int
	Zone     Zone
}
