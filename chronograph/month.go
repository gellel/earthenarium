package chronograph

import (
	"fmt"
	"time"
)

type Month struct {
	Literal time.Month
	Name    string
	Number  int
}

func (month *Month) String() string {
	return fmt.Sprintf("{%s, %v}", month.Name, month.Number)
}
