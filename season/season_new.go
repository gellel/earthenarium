package season

import (
	"github.com/gellel/earthenarium/chronograph"
	"github.com/gellel/earthenarium/hemisphere"
)

func NewSeason(t *chronograph.Time, h *hemisphere.Hemisphere) *Season {
	if h.Latitude.Y == 1 {
		return &Season{
			Term: North.Fetch(t.Month.Literal)}
	}
	return &Season{
		Term: South.Fetch(t.Month.Literal)}
}
