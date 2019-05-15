package climate

import (
	"github.com/gellel/earthenarium/hemisphere"
)

type Climate map[hemisphere.Coordinate]Seasons

func (climate Climate) Fetch(coordinate hemisphere.Coordinate) Seasons {
	seasons, _ := climate.Get(coordinate)
	return seasons
}

func (climate Climate) Get(coordinate hemisphere.Coordinate) (Seasons, bool) {
	seasons, ok := climate[coordinate]
	return seasons, ok
}
