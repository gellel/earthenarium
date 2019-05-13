package hemisphere

func NewZone(key string, area Area) *Zone {
	return &Zone{key: area}
}
