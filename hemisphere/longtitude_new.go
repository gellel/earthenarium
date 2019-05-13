package hemisphere

import (
	"github.com/gellel/earthenarium/longtitude"
)

func NewLongtitude(longtitude *longtitude.Longtitude) Longtitude {
	return Longtitude(longtitude.Value())
}
