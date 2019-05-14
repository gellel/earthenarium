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
		time.December:  NewTerm(hemisphere.N, Winter, q1),
		time.January:   NewTerm(hemisphere.N, Winter, q1),
		time.February:  NewTerm(hemisphere.N, Winter, q1),
		time.March:     NewTerm(hemisphere.N, Spring, q2),
		time.April:     NewTerm(hemisphere.N, Spring, q2),
		time.May:       NewTerm(hemisphere.N, Spring, q2),
		time.June:      NewTerm(hemisphere.N, Summer, q3),
		time.July:      NewTerm(hemisphere.N, Summer, q3),
		time.August:    NewTerm(hemisphere.N, Summer, q3),
		time.September: NewTerm(hemisphere.N, Autumn, q4),
		time.October:   NewTerm(hemisphere.N, Autumn, q4),
		time.November:  NewTerm(hemisphere.N, Autumn, q4)}
)
var (
	South = &Calendar{
		time.December:  NewTerm(hemisphere.S, Summer, q1),
		time.January:   NewTerm(hemisphere.S, Summer, q1),
		time.February:  NewTerm(hemisphere.S, Summer, q1),
		time.March:     NewTerm(hemisphere.S, Autumn, q2),
		time.April:     NewTerm(hemisphere.S, Autumn, q2),
		time.May:       NewTerm(hemisphere.S, Autumn, q2),
		time.June:      NewTerm(hemisphere.S, Winter, q3),
		time.July:      NewTerm(hemisphere.S, Winter, q3),
		time.August:    NewTerm(hemisphere.S, Winter, q3),
		time.September: NewTerm(hemisphere.S, Spring, q4),
		time.October:   NewTerm(hemisphere.S, Spring, q4),
		time.November:  NewTerm(hemisphere.S, Spring, q4)}
)

var (
	Hemisphere = &Seasonality{
		North: North,
		South: South}
)
