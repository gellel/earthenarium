package season

import (
	"github.com/gellel/earthenarium/chronograph"
)

type Season struct {
	Begins *chronograph.Time
	Ends   *chronograph.Time
	Term   *Term
}
