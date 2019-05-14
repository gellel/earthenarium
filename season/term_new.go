package season

func NewTerm(hemisphere, name string, months *Months) *Term {
	return &Term{
		Hemisphere: hemisphere,
		Months:     months,
		Name:       name}
}
