package humidity

import (
	"math/rand"
	"time"

	"github.com/gellel/earthenarium/hemisphere"
)

func NewHumidity(h *hemisphere.Hemisphere) float32 {
	rand.Seed(time.Now().UnixNano())
	switch h.Latitude.Name {
	case hemisphere.CapricornLabel, hemisphere.CancerLabel:
		max, min := float32(80), float32(30)
		return min + rand.Float32()*(max-min)
	case hemisphere.AntarcticLabel, hemisphere.ArcticLabel:
		max, min := float32(67.5), float32(20)
		return min + rand.Float32()*(max-min)
	}
	max, min := float32(50), float32(85)
	return min + rand.Float32()*(max-min)

}
