package hemisphere

import (
	"github.com/gellel/earthenarium/longtitude"
)

func NewLongtitude(longtitude *longtitude.Longtitude) Longtitude {
	l := longtitude.Value()
	if l > prime && l <= east {
		return Longtitude{East, eastLabel, eastLabel}
	}
	if l < prime && l >= west {
		return Longtitude{West, westLabel, westLabel}
	}
	return Longtitude{Prime, primeLabel, primeLabel}
}
