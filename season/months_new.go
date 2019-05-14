package season

import "time"

func NewMonths(months ...time.Month) *Months {
	return &Months{months[0], months[1], months[2]}
}
