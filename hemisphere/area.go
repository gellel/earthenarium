package hemisphere

type Area [3]Coordinate

func (area Area) Peek(i int) Coordinate {
	return area[i]
}
