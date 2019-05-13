package hemisphere

func NewArea(max, mean, min float32) Area {
	return Area{
		NewCoordinate(max),
		NewCoordinate(mean),
		NewCoordinate(min)}
}
