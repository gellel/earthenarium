package chronograph_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/chronograph"
)

func TestNewLocation(testing *testing.T) {
	zone := chronograph.NewZone("America")
	city := chronograph.NewCity("New York")
	location := chronograph.NewLocation(zone, city)
	fmt.Println(location)
}
