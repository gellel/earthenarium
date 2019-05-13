package hemisphere

import (
	"fmt"

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
	if l > prime {
		return Longtitude{Prime, primeLabel, fmt.Sprintf("%s east", primeLabel)}
	}
	return Longtitude{Prime, primeLabel, fmt.Sprintf("%s west", primeLabel)}
}
