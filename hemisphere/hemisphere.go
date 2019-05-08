package hemisphere

const (
	Eastern  Hemisphere = Hemisphere("Eastern")
	Northern Hemisphere = Hemisphere("Northern")
	Southern Hemisphere = Hemisphere("Southern")
	Western  Hemisphere = Hemisphere("Western")
)

type Hemisphere string

func (hemisphere Hemisphere) Latitude() bool {
	return hemisphere == Northern || hemisphere == Southern
}

func (hemisphere Hemisphere) String() string {
	return string(hemisphere)
}
