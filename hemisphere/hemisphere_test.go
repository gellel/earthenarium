package hemisphere_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/hemisphere"

	"github.com/gellel/earthenarium/latitude"
	"github.com/gellel/earthenarium/longtitude"
)

func TestNewHemisphere(testing *testing.T) {
	lat := latitude.Random()
	long := longtitude.Random()
	hem := hemisphere.NewHemisphere(lat, long)
	fmt.Println(hem.Latitude.Name())
}
