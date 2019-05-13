package latitude

const (
	errorLatitude string = "latitude \"%v\" exceeds supported range; cannot eclipse \"%v\" or \"%v\""
)

const (
	Maximum float32 = 90
	Minimum float32 = -Maximum
	Mean    float32 = Maximum + Minimum
)

const (
	Parallel  float32 = 1
	Parallels float32 = ((Maximum - Minimum) * float32(Parallel))
)

const (
	Measurement string = "Degrees"
)
