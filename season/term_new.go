package season

func NewTerm(name string, months *Months) *Term {
	return &Term{
		Months: months,
		Name:   name}
}
