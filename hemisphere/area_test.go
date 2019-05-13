package hemisphere_test

import (
	"testing"

	"github.com/gellel/earthenarium/hemisphere"
)

func TestNewArea(testing *testing.T) {
	hemisphere.Southern.Each(func(key string, area hemisphere.Area) {
		if ok := area[0].Less(area[1]); ok != true {
			testing.Fatalf("area[0] (%v) cannot be greater than area[1] (%v) for southern zone", area[0], area[1])
		}
		if ok := ((area[0] + area[2]) / 2) == area[1]; ok != true {
			testing.Fatalf("area[1] (%v) must be an equal division of a[0] and a[2]", area[1])
		}
	})
	hemisphere.Northern.Each(func(key string, area hemisphere.Area) {
		if ok := area[0].More(area[1]); ok != true {
			testing.Fatalf("area[0] (%v) cannot be less than area[1] (%v) for southern zone", area[0], area[1])
		}
		if ok := ((area[0] + area[2]) / 2) == area[1]; ok != true {
			testing.Fatalf("area[1] (%v) must be an equal division of a[0] and a[2]", area[1])
		}
	})
	hemisphere.Center.Each(func(key string, area hemisphere.Area) {
		if ok := area[0].Absolute() == area[2].Absolute(); ok != true {
			testing.Fatalf("area[0] (%v) and area[2] (%v) must be equal when cast absolutely", area[0], area[2])
		}
		if ok := area[1] == 0; ok != true {
			testing.Fatalf("area[1] must be zero")
		}
	})
}
