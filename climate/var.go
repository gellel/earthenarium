package climate

import (
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/season"
)

var (
	Arctic = Climate{
		// polar
		hemisphere.Arctic[0]: {
			season.Winter: -15,
			season.Spring: -10,
			season.Summer: -5,
			season.Autumn: -10},
		// tundra
		hemisphere.Arctic[1]: {
			season.Winter: -9,
			season.Spring: -2,
			season.Summer: 1,
			season.Autumn: -1},
		// temperate
		hemisphere.Arctic[2]: {
			season.Winter: -2,
			season.Spring: 9,
			season.Summer: 15,
			season.Autumn: 6}}
)
