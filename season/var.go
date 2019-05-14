package season

import (
	"time"

	"github.com/gellel/earthenarium/hemisphere"
)

var (
	q1 = NewMonths(
		time.December,
		time.January,
		time.February)
)
var (
	q2 = NewMonths(
		time.March,
		time.April,
		time.May)
)
var (
	q3 = NewMonths(
		time.June,
		time.July,
		time.August)
)
var (
	q4 = NewMonths(
		time.September,
		time.October,
		time.November)
)
var (
	North = &Calendar{
		time.December:  NewTerm(hemisphere.N, winter, q1),
		time.January:   NewTerm(hemisphere.N, winter, q1),
		time.February:  NewTerm(hemisphere.N, winter, q1),
		time.March:     NewTerm(hemisphere.N, spring, q2),
		time.April:     NewTerm(hemisphere.N, spring, q2),
		time.May:       NewTerm(hemisphere.N, spring, q2),
		time.June:      NewTerm(hemisphere.N, summer, q3),
		time.July:      NewTerm(hemisphere.N, summer, q3),
		time.August:    NewTerm(hemisphere.N, summer, q3),
		time.September: NewTerm(hemisphere.N, autumn, q4),
		time.October:   NewTerm(hemisphere.N, autumn, q4),
		time.November:  NewTerm(hemisphere.N, autumn, q4)}
)
var (
	South = &Calendar{
		time.December:  NewTerm(hemisphere.S, summer, q1),
		time.January:   NewTerm(hemisphere.S, summer, q1),
		time.February:  NewTerm(hemisphere.S, summer, q1),
		time.March:     NewTerm(hemisphere.S, autumn, q2),
		time.April:     NewTerm(hemisphere.S, autumn, q2),
		time.May:       NewTerm(hemisphere.S, autumn, q2),
		time.June:      NewTerm(hemisphere.S, winter, q3),
		time.July:      NewTerm(hemisphere.S, winter, q3),
		time.August:    NewTerm(hemisphere.S, winter, q3),
		time.September: NewTerm(hemisphere.S, spring, q4),
		time.October:   NewTerm(hemisphere.S, spring, q4),
		time.November:  NewTerm(hemisphere.S, spring, q4)}
)

var (
	Hemisphere = &Seasonality{
		North: North,
		South: South}
)
