package chronograph_test

import (
	"testing"
	"time"

	"github.com/gellel/earthenarium/chronograph"
)

func TestNewMonth(testing *testing.T) {
	chronograph.NewMonth(time.March)
}
