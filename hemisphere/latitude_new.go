package hemisphere

import (
	"fmt"

	"github.com/gellel/earthenarium/latitude"
)

func NewLatitude(latitude *latitude.Latitude) *Latitude {
	l := latitude.Value()
	if l >= arctic && l <= north {
		return &Latitude{
			Area:      &Arctic,
			Name:      arcticLabel,
			Reference: northLabel,
			Y:         1}
	}
	if l >= cancer && l < arctic {
		return &Latitude{
			Area:      &Cancer,
			Name:      cancerLabel,
			Reference: northLabel,
			Y:         1}
	}
	if l <= capricorn && l > antarctic {
		return &Latitude{
			Area:      &Capricorn,
			Name:      capricornLabel,
			Reference: southLabel,
			Y:         -1}
	}
	if l <= antarctic && l >= south {
		return &Latitude{
			Area:      &Antarctic,
			Name:      antarcticLabel,
			Reference: southLabel,
			Y:         -1}
	}
	if l > equator {
		return &Latitude{
			Area:      &Equator,
			Name:      equatorLabel,
			Reference: fmt.Sprintf("%s north", equatorLabel),
			Y:         1}
	}
	return &Latitude{
		Area:      &Equator,
		Name:      equatorLabel,
		Reference: fmt.Sprintf("%s south", equatorLabel),
		Y:         -1}
}
