package temperature

type Temperatures []*Temperature

func (temperatures *Temperatures) Average() float32 {
	s := float32(0)
	for _, t := range *temperatures {
		s = s + t.Value()
	}
	return s / float32(temperatures.Len())
}

func (temperatures *Temperatures) Len() int {
	return len(*temperatures)
}
