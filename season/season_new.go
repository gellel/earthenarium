package season

import (
	"time"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/hemisphere"
)

func NewSeason(t *chronograph.Time, h *hemisphere.Hemisphere) *Season {

	calendar := Hemisphere.Fetch(h)
	term := calendar.Fetch(t.Month.Literal)

	switch t.Month.Literal {
	case time.December, time.January, time.February:
		first := time.Date(t.Year.Number-1, term.Months[0], 1, 0, 0, 0, 0, t.Location)
		begins := chronograph.NewTime(&first)
		last := time.Date(begins.Year.Number+1, term.Months[2], 1, 0, 0, 0, 0, t.Location).AddDate(0, 1, 0).Add(time.Nanosecond * -1)
		ends := chronograph.NewTime(&last)
		return &Season{
			Begins: begins,
			Ends:   ends,
			Term:   term}
	}
	first := time.Date(t.Year.Number, term.Months[0], 1, 0, 0, 0, 0, t.Location)
	begins := chronograph.NewTime(&first)
	last := time.Date(begins.Year.Number, term.Months[2], 1, 0, 0, 0, 0, t.Location).AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	ends := chronograph.NewTime(&last)
	return &Season{
		Begins: begins,
		Ends:   ends,
		Term:   term}
}
