package season

import (
	"time"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/hemisphere"
)

func NewSeason(t *chronograph.Time, h *hemisphere.Hemisphere) *Season {
	calendar := Hemisphere.Fetch(h)
	term := calendar.Fetch(t.Month.Literal)
	yearStart := t.Year.Number
	yearEnd := t.Year.Number
	month := t.Month.Literal
	if month == time.December || month == time.January || month == time.February {
		yearStart = yearStart - 1
	}
	begins := chronograph.NewTime(time.Date(yearStart, term.Months[0], 1, 0, 0, 0, 0, t.Location))
	ends := chronograph.NewTime(time.Date(yearEnd, (term.Months[2] + 1), 1, 0, 0, 0, 0, t.Location).Add(time.Nanosecond * -1))
	return &Season{
		Begins:     begins,
		Ends:       ends,
		Hemisphere: term.Hemisphere,
		Months:     term.Months,
		Name:       term.Name}
}
