package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
)

func NewLatitude(latitude *latitude.Latitude) *Latitude {
	l := latitude.Value()
	if l >= arctic && l <= north {
		return &Latitude{
			Area:      &Arctic,
			Name:      ArcticLabel,
			Reference: NorthLabel,
			Y:         1}
	}
	if l >= cancer && l < arctic {
		return &Latitude{
			Area:      &Cancer,
			Name:      CancerLabel,
			Reference: NorthLabel,
			Y:         1}
	}
	if l <= capricorn && l > antarctic {
		return &Latitude{
			Area:      &Capricorn,
			Name:      CapricornLabel,
			Reference: SouthLabel,
			Y:         -1}
	}
	if l <= antarctic && l >= south {
		return &Latitude{
			Area:      &Antarctic,
			Name:      AntarcticLabel,
			Reference: SouthLabel,
			Y:         -1}
	}
	if l > equator {
		return &Latitude{
			Area:      &Equator,
			Name:      EquatorLabel,
			Reference: NorthLabel,
			Y:         1}
	}
	return &Latitude{
		Area:      &Equator,
		Name:      EquatorLabel,
		Reference: SouthLabel,
		Y:         -1}
}
