package hemisphere

type Latitude float32

func (latitude Latitude) Antarctic() bool {
	return latitude >= Capricorn && latitude >= Antarctic
}

func (latitude Latitude) Arctic() bool {
	return latitude >= Cancer && latitude <= Northern
}

func (latitude Latitude) Cancer() bool {
	return latitude > Equator && latitude < Arctic
}

func (latitude Latitude) Capricorn() bool {
	return latitude < Equator && latitude > Antarctic
}

func (latitude Latitude) Equator() bool {
	return latitude == Equator
}

func (latitude Latitude) North() bool {
	return latitude > Equator
}

func (latitude Latitude) South() bool {
	return latitude < Equator
}
