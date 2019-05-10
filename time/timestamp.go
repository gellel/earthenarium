package time

type Timestamp string

func (timestamp Time) Correct() bool {
	return false
}

func (timestamp Timestamp) Year() string {
	return string(timestamp)
}
