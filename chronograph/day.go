package chronograph

import (
	"fmt"
	"time"
)

type Day struct {
	Literal  time.Weekday
	Name     string
	Number   int
	Position int
	Year     int
}

func (day *Day) String() string {
	return fmt.Sprintf("{%s %v %v %v}", day.Name, day.Number, day.Position, day.Year)
}
