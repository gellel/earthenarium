package chronograph

import (
	"regexp"
	"strings"
)

// NewTimestamp instantiates a new Timestamp pointer.
func NewTimestamp(time string) *Timestamp {
	reg := regexp.MustCompile("[^TZ\\-0-9:.]")
	time = strings.ToUpper(time)
	time = reg.ReplaceAllString(time, "")
	timestamp := Timestamp(time)
	return &timestamp
}
