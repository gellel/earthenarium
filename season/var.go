package season

import "time"

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
		time.December:  NewTerm(winter, q1),
		time.January:   NewTerm(winter, q1),
		time.February:  NewTerm(winter, q1),
		time.March:     NewTerm(spring, q2),
		time.April:     NewTerm(spring, q2),
		time.May:       NewTerm(spring, q2),
		time.June:      NewTerm(summer, q3),
		time.July:      NewTerm(summer, q3),
		time.August:    NewTerm(summer, q3),
		time.September: NewTerm(autumn, q4),
		time.October:   NewTerm(autumn, q4),
		time.November:  NewTerm(autumn, q4)}
)
var (
	South = &Calendar{
		time.December:  NewTerm(summer, q1),
		time.January:   NewTerm(summer, q1),
		time.February:  NewTerm(summer, q1),
		time.March:     NewTerm(autumn, q2),
		time.April:     NewTerm(autumn, q2),
		time.May:       NewTerm(autumn, q2),
		time.June:      NewTerm(winter, q3),
		time.July:      NewTerm(winter, q3),
		time.August:    NewTerm(winter, q3),
		time.September: NewTerm(spring, q4),
		time.October:   NewTerm(spring, q4),
		time.November:  NewTerm(spring, q4)}
)
