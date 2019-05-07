package season

import "github.com/gellel/earthenarium/chronograph"

type Autumn string
type Spring string
type Summer string
type Winter string

type Season struct {
	Begins *chronograph.Time
	Ends   *chronograph.Time
	Name   string
	Spans  *chronograph.Span
}

func NewSeason(hemisphere string, begins, ends *chronograph.Time, spans *chronograph.Span) *Season {
	return &Season{
		Begins: begins,
		Ends:   ends,
		Spans:  spans}
}

func GetSeason(t *chronograph.Time) string {
	return ""
}
