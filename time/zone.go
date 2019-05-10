package time

// Zone contains a fragment.
type Zone string

// String converts a Zone to a go.string.
func (zone Zone) String() string {
	return string(zone)
}
