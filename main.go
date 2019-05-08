package main

import (
	"fmt"

	"github.com/gellel/earthenarium/chronograph"
)

func main() {

	y := "2015-02-28T23:59:59.000Z"

	fmt.Println(chronograph.NewTimeFromISO(y))

}
