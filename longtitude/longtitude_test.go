package longtitude_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/longtitude"
)

func TestNewLongtitude(testing *testing.T) {
	long := longtitude.NewLongtitude(180)
	fmt.Println(long)
}
