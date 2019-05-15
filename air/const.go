package air

var (
	pressure  = 0
	density   = 1
	gravity   = 9.80665   // Gravitational constant (m/s/s)
	gas       = 8.31432   // Universal gas constant (N*m/(mol*K))
	molar     = 0.0289644 // Molar mass of earth's air (kg/mol)
	g         = ((gravity * molar) / gas)
	logarithm = 2.718
)
