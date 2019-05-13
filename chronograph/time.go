package chronograph

import "time"

type Time struct {
	Day        *Day
	Location   *time.Location
	Minute     int
	Month      *Month
	Nanosecond int
	Second     int
	Time       *time.Time
	Unix       int64
	Year       int
}
