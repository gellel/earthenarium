package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
)

func NewLatitude(latitude *latitude.Latitude) Latitude {
	return Latitude(latitude.Value())
}
