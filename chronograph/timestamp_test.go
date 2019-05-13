package chronograph_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/chronograph"
)

func TestNewTimestamp(testing *testing.T) {
	timestamp := chronograph.NewTimestamp(chronograph.Layout)
	fmt.Println(timestamp)
	fmt.Println(timestamp.Year(), timestamp.Month(), timestamp.Day())
	fmt.Println(timestamp.Hour(), timestamp.Minute(), timestamp.Second())
	fmt.Println(timestamp.Correct())
	fmt.Println(timestamp.Parse())
}
