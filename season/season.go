package season

import (
	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/coordinate"
	"github.com/gellel/earthenarium/hemisphere"
)

type Season struct {
	Begins *chronograph.Time
	Ends   *chronograph.Time
	Name   string
	Spans  *chronograph.Span
}

func NewSeason(latlong *coordinate.Coordinate) *Season {
	return &Season{}
}

func GetSeason(latlong *coordinate.Coordinate, t *chronograph.Time) string {
	switch latlong.Latitude.Region() {
	case hemisphere.Antarctic:
	case hemisphere.Arctic:
	case hemisphere.Northern:
	case hemisphere.Southern:
	}
	return ""
}
