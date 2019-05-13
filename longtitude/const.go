package longtitude

const (
	errorLongtitude string = "latitude \"%v\" exceeds supported range; cannot eclipse \"%v\" or \"%v\""
)

const (
	Maximum float32 = 180
	Minimum float32 = -Maximum
	Mean    float32 = (Maximum + Minimum)
)

const (
	Meridian  float32 = 1
	Meridians float32 = ((Maximum - Minimum) * float32(Meridian))
)

const (
	Measurement string = "Degrees"
)
