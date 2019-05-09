package season

import (
	"fmt"
	"time"

	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/hemisphere"
)

const (
	// TimestampStart expresses the lower datetime bounds of a seasonal timestamp.
	timestampStart string = "T00:00:00.000Z"
	// TimestampEnd expresses the upper datetime bounds of a seasonal timestamp.
	timestampEnd string = "T23:59:59.000Z"
)

const (
	// Autumn is the name expressing the season of Autumn.
	Autumn string = "Autumn"
	// Spring is the name expressing the season of Spring.
	Spring string = "Spring"
	// Summer is the name expressing the season of Summer.
	Summer string = "Summer"
	// Winter is the name expressing the season of Winter.
	Winter string = "Winter"
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

func format(a, b int, timestamps []string) (string, string) {
	return fmt.Sprintf(timestamps[0], a), fmt.Sprintf(timestamps[1], b)
}

// Northernize localises season conditions to align to the chronological timings of the Earth's northern hemisphere.
func Northernize(t *chronograph.Time) string {
	switch t.Month.T {
	case time.December, time.January, time.February:
		return Winter
	case time.March, time.April, time.May:
		return Spring
	case time.June, time.July, time.August:
		return Summer
	default:
		return Autumn
	}
}

// Select automatically determines the appropriate seasonal condition based on the regional input string and chronograph.Time struct.
func Select(region string, t *chronograph.Time) string {
	ok := region == hemisphere.Northern || region == hemisphere.Southern
	if ok != true {
		panic(region)
	}
	if region == hemisphere.Northern {
		return Northernize(t)
	}
	return Southernize(t)
}

// Southernize localises season conditions to align to the chronological timings of the Earth's southern hemisphere.
func Southernize(t *chronograph.Time) string {
	switch t.Month.T {
	case time.December, time.January, time.February:
		return Summer
	case time.March, time.April, time.May:
		return Autumn
	case time.June, time.July, time.August:
		return Winter
	default:
		return Spring
	}
}

// Timespan breaks down the datetime data held by the chronograph.Time struct to generate a lower and upper bounds of a season.
// Season span generated is not regionalized and needs to be cross-checked against a hemisphere.
func Timespan(t *chronograph.Time) (string, string) {
	switch t.Month.T {
	case time.December, time.January, time.February:
		if t.Month.T == time.December {
			return format(t.Year, (t.Year + 1), T12X02)
		}
		return format((t.Year - 1), t.Year, T12X02)
	case time.March, time.April, time.May:
		return format(t.Year, t.Year, T03X05)
	case time.June, time.July, time.August:
		return format(t.Year, t.Year, T06X08)
	default:
		return format(t.Year, t.Year, T09X11)
	}
}

// NewSeason creates a new Season pointer.
func NewSeason(region string, t *chronograph.Time) *Season {
	ok := region == hemisphere.Northern || region == hemisphere.Southern
	if ok != true {
		panic(region)
	}
	name := Select(region, t)
	a, b := Timespan(t)
	timeBegins := chronograph.NewTimeFromISO(a)
	timeEnds := chronograph.NewTimeFromISO(b)
	timeSpan := chronograph.NewSpan(timeBegins.T, timeEnds.T)
	return &Season{
		Begins:     timeBegins,
		Ends:       timeEnds,
		Hemisphere: region,
		Name:       name,
		Spans:      timeSpan}
}

// NewSeasonNorth creates a new Season pointer oriented for northern hemisphere climates.
func NewSeasonNorth(t *chronograph.Time) *Season {
	return NewSeason(hemisphere.Northern, t)
}

// NewSeasonSouth creates a new Season pointer oriented for southern hemisphere climates.
func NewSeasonSouth(t *chronograph.Time) *Season {
	return NewSeason(hemisphere.Southern, t)
}

// Season stores a chronological data that expresses the Earth-like season.
type Season struct {
	Begins     *chronograph.Time
	Ends       *chronograph.Time
	Hemisphere string
	Name       string
	Spans      *chronograph.Span
}

// Previous updates the Season pointer to its previous chronological iteration.
func (season *Season) Previous() *Season {
	*season = *NewSeason(season.Hemisphere, season.Begins.AddDays(-1))
	return season
}

// Next updates the Season pointer to its next chronological iteration.
func (season *Season) Next() *Season {
	*season = *NewSeason(season.Hemisphere, season.Ends.AddDays(1))
	return season
}
