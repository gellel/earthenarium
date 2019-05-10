package time_test

import (
	"fmt"
	"testing"
	t "time"

	"github.com/gellel/earthenarium/time"
)

func TestNewMonth(testing *testing.T) {
	const (
		f string = "{name: %s, number: %v}"
		e string = "time.NewDay(time.Now()) did not return " + f + ". returned " + f + " instead."
	)
	var (
		now    = t.Now()
		name   = now.Month().String()
		number = int(now.Month())
	)
	var (
		month = time.NewMonth(&now)
	)
	if month.Name != name || month.Number != number {
		testing.Fatal(fmt.Errorf(e, name, number, month.Name, month.Number))
	}
}
