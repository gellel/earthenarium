package climate

import (
	"github.com/gellel/earthenarium/hemisphere"
)

type Climate map[hemisphere.Coordinate]map[string]float32

func (climate Climate) Fetch(coordinate hemisphere.Coordinate) map[string]float32 {
	seasons, _ := climate.Get(coordinate)
	return seasons
}
func (climate Climate) Get(coordinate hemisphere.Coordinate) (map[string]float32, bool) {
	seasons, ok := climate[coordinate]
	return seasons, ok
}
