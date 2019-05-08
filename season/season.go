package season

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"
)

const (
	start string = "T00:00:00Z"
	end   string = "T23:59:59Z"
)

var (
	// T12X02 expresses the epoch timestamp from December to Feburary of the next year.
	T12X02 = []string{
		("%v-12-01" + start),
		("%v-02-28" + end)}

	// T03X05 expresses the epoch timestamp from March to May of the current year.
	T03X05 = []string{
		("%v-03-01" + start),
		("%v-05-31" + end)}

	// T06X08 expresses the epoch timestamp from June to August of the current year.
	T06X08 = []string{
		("%v-06-01" + start),
		("%v-08-31" + end)}

	// T09X11 expresses epoch timestamp from September to November of the current year.
	T09X11 = []string{
		("%v-09-01" + start),
		("%v-11-30" + end)}
)

var (
	// Epochs contains available date ranges that can be computed for a calendar year.
	Epochs = [][]string{
		T12X02, T03X05, T06X08, T09X11}
)

func Seasonalize(t *chronograph.Time) (string, string) {
	switch t.Month.Name {
	case "December", "Janurary", "Feburary":
		if t.Month.Name == "December" {
			return fmt.Sprintf(T12X02[0], t.Year), fmt.Sprintf(T12X02[1], t.Year+1)
		}
		return fmt.Sprintf(T12X02[0], t.Year), fmt.Sprintf(T12X02[1], t.Year)
	}
	return "", ""
}

func SeasonFromTime(t *chronograph.Time) {

}
