package hemisphere

import "strings"

type Label string

func (label Label) Key() string {
	return strings.Replace(string(label), " ", "_", -1)
}

func (label Label) String() string {
	return string(label)
}
