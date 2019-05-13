package chronograph

import "time"

// NewDay instantiates a new Day pointer.
func NewDay(day, days int, weekday time.Weekday) *Day {
	return &Day{
		Name:     weekday.String(),
		Number:   day,
		Position: int(weekday),
		Year:     days}
}
