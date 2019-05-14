package temperature

func NewTemperature(temperature float32) *Temperature {
	t := Temperature(temperature)
	return &t
}
