package calendar_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/calendar"
)

func TestDay(testing *testing.T) {

	day := calendar.NewDay()
	fmt.Println(day.Hour())
}
