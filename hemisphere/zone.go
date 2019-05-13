package hemisphere

type Zone map[string]Area

func (zone *Zone) Add(key string, area Area) {
	(*zone)[key] = area
}

func (zone *Zone) Each(f func(key string, area Area)) {
	for key, value := range *zone {
		f(key, value)
	}
}

func (zone *Zone) Concat(zones ...*Zone) *Zone {
	for _, z := range zones {
		for key, value := range *z {
			zone.Add(key, value)
		}
	}
	return zone
}
