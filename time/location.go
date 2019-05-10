package time

import (
	"regexp"
	"strings"
)

// Location contains the string value required for Go time.LoadLocation to fetch timezone data.
type Location string

func (location Location) Correct() bool {
	return strings.Contains(string(location), "/") && (!regexp.MustCompile("\\d|\\s").MatchString(string(location)))
}

func (location Location) String() string {
	return string(location)
}
