package air

type Pressure float32

func (pressure Pressure) Hpa() float32 {
	return float32(pressure / 100)
}
