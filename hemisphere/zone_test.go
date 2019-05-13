package hemisphere_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/hemisphere"
)

func TestNewHemisphere(testing *testing.T) {
	a := hemisphere.NewZone("a", hemisphere.NewArea(1, 0, -1))
	b := hemisphere.NewZone("b", hemisphere.NewArea(-1, 0, 1))
	c := (&hemisphere.Zone{}).Concat(a, b)
	fmt.Println(c)
}
