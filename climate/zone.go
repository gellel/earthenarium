package climate

type Zone map[float32]Climate

func (zone Zone) Fetch(key float32) Climate {
	climate, _ := zone.Get(key)
	return climate
}

func (zone Zone) Get(key float32) (Climate, bool) {
	climate, ok := zone[key]
	return climate, ok
}
