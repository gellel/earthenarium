package chronograph

import "errors"

type Zone string

func (zone *Zone) Correct() (ok bool) {
	return !ok
}

func (zone *Zone) Error() error {
	return errors.New("")
}

func (zone *Zone) String() string {
	return string(*zone)
}
