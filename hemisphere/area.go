package hemisphere

import (
	"github.com/gellel/earthenarium/latitude"
)

type Area [3]Coordinate

func (area Area) Lat(latitude *latitude.Latitude) float32 {
	if latitude.Value() > area[1].Value() {
		if latitude.Value() > float32(area[1]+area[2])/2 {
			return area[2].Value()
		}
	}
	if latitude.Value() < area[1].Value() {
		if latitude.Value() < float32(area[0]+area[1])/2 {
			return area[0].Value()
		}
	}
	return area[1].Value()
}
