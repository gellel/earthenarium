package climate

type Seasons map[string]float32

func (seasons Seasons) Fetch(key string) float32 {
	temperature, _ := seasons.Get(key)
	return temperature
}

func (seasons Seasons) Get(key string) (float32, bool) {
	temperature, ok := seasons[key]
	return temperature, ok
}
