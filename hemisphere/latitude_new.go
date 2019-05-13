package hemisphere

import (
	"fmt"

	"github.com/gellel/earthenarium/latitude"
)

func NewLatitude(latitude *latitude.Latitude) Latitude {
	l := latitude.Value()
	if l >= arctic && l <= north {
		return Latitude{Arctic, arcticLabel, northLabel}
	}
	if l >= cancer && l < arctic {
		return Latitude{Cancer, cancerLabel, northLabel}
	}
	if l <= capricorn && l > antarctic {
		return Latitude{Capricorn, capricornLabel, southLabel}
	}
	if l <= antarctic && l >= south {
		return Latitude{Antarctic, antarcticLabel, southLabel}
	}
	if l > equator {
		return Latitude{Equator, equatorLabel, fmt.Sprintf("%s north", equatorLabel)}
	}
	return Latitude{Equator, equatorLabel, fmt.Sprintf("%s south", equatorLabel)}
}
