package chronograph

import "fmt"

// NewTimespan instantiates a new Timespan pointer.
func NewTimespan(begins, ends *Time) *Timespan {
	if ok := begins.Unix < ends.Unix; !ok {
		panic(fmt.Errorf("%v > %v", begins.Unix, ends.Unix))
	}
	seconds := ends.Time.Sub(*begins.Time).Seconds()
	minutes := (seconds / 60)
	hours := (minutes / 60)
	days := (hours / 24)
	months := (ends.Month.Number - begins.Month.Number)
	years := (ends.Year - begins.Year)
	return &Timespan{
		Days:    int(days),
		Hours:   int(hours),
		Minutes: int(minutes),
		Months:  months,
		Seconds: int(seconds),
		Years:   years}
}
