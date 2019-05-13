package temperature

type Temperature float32

func (temperature *Temperature) Value() float32 {
	return float32(*temperature)
}

func (temperature *Temperature) Fahrenheit() float32 {
	return ((float32(*temperature) * 9 / 5) + Fahrenheit)
}

func (temperature *Temperature) Kelvin() float32 {
	return (float32(*temperature) * Kelvin)
}
