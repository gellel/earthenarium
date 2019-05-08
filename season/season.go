package season

import (
	"fmt"
	"time"

	"github.com/gellel/earthenarium/chronograph"
)

const (
	// TimestampStart expresses the lower datetime bounds of a seasonal timestamp.
	timestampStart string = "T00:00:00.000Z"
	// TimestampEnd expresses the upper datetime bounds of a seasonal timestamp.
	timestampEnd string = "T23:59:59.000Z"
)

var (
	// T12X02 expresses the epoch timestamp from December to Feburary of the next year.
	T12X02 = []string{
		("%v-12-01" + timestampStart),
		("%v-02-28" + timestampEnd)}
	// T03X05 expresses the epoch timestamp from March to May of the current year.
	T03X05 = []string{
		("%v-03-01" + timestampStart),
		("%v-05-31" + timestampEnd)}
	// T06X08 expresses the epoch timestamp from June to August of the current year.
	T06X08 = []string{
		("%v-06-01" + timestampStart),
		("%v-08-31" + timestampEnd)}
	// T09X11 expresses epoch timestamp from September to November of the current year.
	T09X11 = []string{
		("%v-09-01" + timestampStart),
		("%v-11-30" + timestampEnd)}
	// Timestamps contains available date ranges that can be computed for a calendar year.
	// Slices are ordered chronologically from December of the previous year to November
	// of the previous year + 1.
	Timestamps = [][]string{
		T12X02, T03X05, T06X08, T09X11}
)

func fetch(a, b int, timestamp []string) (string, string) {
	return fmt.Sprintf(timestamp[0], a), fmt.Sprintf(timestamp[1], b)
}

func timespan(begins, ends string) *Timespan {
	a := chronograph.NewTimeFromISO(begins)
	b := chronograph.NewTimeFromISO(ends)
	s := chronograph.NewSpan(a.T, b.T)
	return &Timespan{
		Begins: a,
		Ends:   b,
		Spans:  s}
}

func NewSeason(region string, timespan *Timespan) *Season {
	return &Season{}
}

func NewTimespan(t *chronograph.Time) *Timespan {
	switch t.Month.T {
	case time.December, time.January, time.February:
		if t.Month.T == time.December {
			return timespan(fetch(t.Year, (t.Year + 1), T12X02))
		}
		return timespan(fetch((t.Year - 1), t.Year, T12X02))
	case time.March, time.April, time.May:
		return timespan(fetch(t.Year, t.Year, T03X05))
	case time.June, time.July, time.August:
		return timespan(fetch(t.Year, t.Year, T06X08))
	default:
		return timespan(fetch(t.Year, t.Year, T09X11))
	}
}

type Season struct {
	Begins     *chronograph.Time
	Ends       *chronograph.Time
	Hemisphere string
	Name       string
	Spans      *chronograph.Span
}

type Timespan struct {
	Begins *chronograph.Time
	Ends   *chronograph.Time
	Spans  *chronograph.Span
}
