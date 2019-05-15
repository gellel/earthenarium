package air

import (
	"math"

	"github.com/gellel/earthenarium/elevation"
	"github.com/gellel/earthenarium/temperature"
)

func calculate(unit int, t *temperature.Temperature, elevation *elevation.Elevation) float64 {

	var (
		h  = elevation.Float64()
		hb float64 // Standard height
		lb float64 // Standard lapse rate
		rb float64 // Standard density
		pb float64 // Standard pressure
		tb float64 // Standard temperature
		c  float64 // interim value (standard density or pressure)
		m  float64 // interim value (mantissa of result)
		e  float64 // interim value (exponent of result)
	)
	// set metrics (Subscript b) from atmospheric layer
	if h < 11000 { // b = 0
		hb = 0
		lb = -0.0065
		pb = 101325.00
		rb = 1.2250
		tb = 288.15
	} else if h < 20000 { // b = 1
		hb = 11000
		lb = 0.0
		pb = 22632.10
		rb = 0.36391
		tb = 216.65
	} else if h < 32000 { // b = 2
		hb = 20000
		lb = 0.001
		pb = 5474.89
		rb = 0.08803
		tb = 216.65
	} else if h < 47000 { // b = 3
		hb = 32000
		lb = 0.0028
		pb = 868.02
		rb = 0.01322
		tb = 228.65
	} else if h < 51000 { // b = 4
		hb = 47000
		lb = 0.0
		pb = 110.91
		rb = 0.00143
		tb = 270.65
	} else if h < 71000 { // b = 5
		hb = 51000
		lb = -0.0028
		pb = 66.94
		rb = 0.00086
		tb = 270.65
	} else { // b = 6
		hb = 71000
		lb = -0.002
		pb = 3.96
		rb = 0.000064
		tb = 214.65
	}
	if unit == pressure {
		c = pb
	} else {
		c = rb
	}
	if lb == 0 {
		m = logarithm
		e = -g * (h - hb) / tb
	} else if unit == pressure {
		m = tb / (tb + lb*(h-hb))
		e = g / lb
	} else if unit == density {
		m = (tb + lb*(h-hb)) / tb
		e = -g/lb - 1
	}
	return c * math.Pow(m, e)
}
