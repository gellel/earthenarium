package time

// City is a expression of a timezone city used for time.LoadLocation.
type City string

func (city City) String() string {
	return string(city)
}
