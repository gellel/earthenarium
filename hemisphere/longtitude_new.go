package hemisphere

import (
	"github.com/gellel/earthenarium/longtitude"
)

func NewLongtitude(longtitude *longtitude.Longtitude) *Longtitude {
	l := longtitude.Value()
	if l > prime && l <= east {
		return &Longtitude{
			Area:      &East,
			Name:      EastLabel,
			Reference: EastLabel,
			X:         1}
	}
	if l < prime && l >= west {
		return &Longtitude{
			Area:      &West,
			Name:      WestLabel,
			Reference: WestLabel,
			X:         -1}
	}
	if l > prime {
		return &Longtitude{
			Area:      &Prime,
			Name:      PrimeLabel,
			Reference: EastLabel,
			X:         1}
	}
	return &Longtitude{
		Area:      &Prime,
		Name:      PrimeLabel,
		Reference: WestLabel,
		X:         -1}
}
