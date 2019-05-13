package chronograph

import "errors"

type Location string

func (location *Location) Correct() (ok bool) {
	return !ok
}

func (location *Location) Error() error {
	return errors.New("")
}

func (location *Location) String() (local string) {
	local = string(*location)
	return local
}
