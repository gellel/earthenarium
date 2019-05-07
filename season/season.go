package season

import (
	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/hemisphere"
)

type Season struct {
	Begins *chronograph.Time
	Ends   *chronograph.Time
	Name   string
	Spans  *chronograph.Span
}

func NewSeason(hemisphere *hemisphere.Hemisphere, begins, ends *chronograph.Time, spans *chronograph.Span) *Season {
	return &Season{
		Begins: begins,
		Ends:   ends,
		Spans:  spans}
}

func GetSeason(t *chronograph.Time) string {
	return ""
}
