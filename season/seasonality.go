package season

import (
	"github.com/gellel/earthenarium/hemisphere"
)

type Seasonality struct {
	North *Calendar
	South *Calendar
}

func (seasonality *Seasonality) Fetch(hemisphere *hemisphere.Hemisphere) *Calendar {
	calendar, _ := seasonality.Get(hemisphere)
	return calendar
}

func (seasonality *Seasonality) Get(hemisphere *hemisphere.Hemisphere) (*Calendar, string) {
	if hemisphere.Latitude.Y == 1 {
		return seasonality.North, hemisphere.Latitude.Reference
	}
	return seasonality.South, hemisphere.Latitude.Reference
}
