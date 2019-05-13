package chronograph_test

import (
	"fmt"
	"testing"

	"github.com/gellel/earthenarium/chronograph"
)

func TestNewCity(testing *testing.T) {
	city := chronograph.NewCity("Addis 2 Ababa")
	fmt.Println(city)
	fmt.Println(city.Correct())
}
