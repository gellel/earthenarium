package chronograph

import "time"

type Time struct {
	Day        *Day
	Hour       int
	Location   *time.Location
	Minute     int
	Month      *Month
	Nanosecond int
	Second     int
	Time       *time.Time
	Unix       int64
	Year       int
}
