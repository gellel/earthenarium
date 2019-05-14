package hemisphere

import (
	"fmt"

	"github.com/gellel/earthenarium/longtitude"
)

func NewLongtitude(longtitude *longtitude.Longtitude) *Longtitude {
	l := longtitude.Value()
	if l > prime && l <= east {
		return &Longtitude{
			Area:      &East,
			Name:      eastLabel,
			Reference: eastLabel,
			X:         1}
	}
	if l < prime && l >= west {
		return &Longtitude{
			Area:      &West,
			Name:      westLabel,
			Reference: westLabel,
			X:         -1}
	}
	if l > prime {
		return &Longtitude{
			Area:      &Prime,
			Name:      primeLabel,
			Reference: fmt.Sprintf("%s east", primeLabel),
			X:         1}
	}
	return &Longtitude{
		Area:      &Prime,
		Name:      primeLabel,
		Reference: fmt.Sprintf("%s west", primeLabel),
		X:         -1}
}
