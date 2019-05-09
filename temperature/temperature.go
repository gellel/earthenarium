package temperature

const (
	Decline     float32 = 6.5
	Measurement string  = "celcius"
)

type Temperature struct {
	Average float32
	Max     float32
	Min     float32
}
