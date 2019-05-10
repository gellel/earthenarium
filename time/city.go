package time

type City string

func (city City) String() string {
	return string(city)
}
