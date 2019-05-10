package time_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/time"
)

func TestNewCity(testing *testing.T) {
	const (
		expect string = "Test_of_City/Goose_Bay"
		input  string = "TEsT __of _ _CiTY_/Goose_Bay"
	)
	const (
		e string = "time.NewCity(\"%s\") did not return \"%s\". returned \"%s\" instead."
	)
	var (
		output string = time.NewCity(input).String()
	)
	if output != expect {
		testing.Fatal(fmt.Errorf(e, input, expect, output))
	}
}
