package hemisphere_test

import (
	"testing"

	"github.com/gellel/earthenarium/longtitude"

	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/latitude"
)

func TestNewCoordinate(testing *testing.T) {
	lat := latitude.Random()
	long := longtitude.Random()
	coordinateLat := hemisphere.NewCoordinateFromLatitude(lat)
	coordinateLong := hemisphere.NewCoordinateFromLongtitude(long)
	if ok := lat.Value() == coordinateLat.Value(); ok != true {
		testing.Fatalf("latitude and coordinate values do not match")
	}
	if ok := long.Value() == coordinateLong.Value(); ok != true {
		testing.Fatalf("longtitude and coordinate values do not match")
	}
}
