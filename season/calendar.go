package season

import "time"

type Calendar map[time.Month]*Term

func (calendar *Calendar) Fetch(month time.Month) *Term {
	term, _ := calendar.Get(month)
	return term
}

func (calendar *Calendar) Get(month time.Month) (*Term, bool) {
	term, ok := (*calendar)[month]
	return term, ok
}
