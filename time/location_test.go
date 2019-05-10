package time_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gellel/earthenarium/time"
)

func TestNewLocation(testing *testing.T) {
	const (
		expect string = "America/New_York"
		input  string = "AmEr1ica/New_ YoRk"
	)
	const (
		e string = "time.NewLocation(\"%s\") did not return \"%s\". returned \"%s\" instead."
	)
	var (
		substrings []string = strings.Split(input, "/")
		output     string   = time.NewLocation(substrings[0], substrings[1]).String()
	)
	if output != expect {
		testing.Fatal(fmt.Errorf(e, input, expect, output))
	}
}
