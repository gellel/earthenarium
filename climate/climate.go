package climate

import (
	"github.com/gellel/earthenarium/hemisphere"
)

type Climate map[hemisphere.Coordinate]map[string]float32
