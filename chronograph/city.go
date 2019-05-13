package chronograph

import "fmt"

type City string

func (city *City) Correct() (ok bool) {
	return !ok
}

func (city *City) Error() error {
	return fmt.Errorf(errorCity, string(*city))
}

func (city *City) String() string {
	return string(*city)
}
