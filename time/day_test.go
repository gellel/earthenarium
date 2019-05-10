package time_test

import (
	"fmt"
	"testing"

	t "time"

	"github.com/gellel/earthenarium/time"
)

func TestNewDay(testing *testing.T) {
	const (
		f string = "{name: %s, number: %v, position: %v}"
		e string = "time.NewDay(time.Now()) did not return " + f + ". returned " + f + " instead."
	)
	var (
		now      = t.Now()
		name     = now.Weekday().String()
		number   = now.Day()
		position = int(now.Weekday())
	)
	var (
		day = time.NewDay(&now)
	)
	if day.Name != name || day.Number != number || day.Position != position {
		testing.Fatal(fmt.Errorf(e, name, number, position, day.Name, day.Number, day.Position))
	}
}
