package latitude_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/latitude"
)

func TestNewLatitude(testing *testing.T) {
	a := latitude.NewLatitude(90)
	b := latitude.NewLatitude(71)
	fmt.Println(a, b)
	fmt.Println(a.Max(b))
}
