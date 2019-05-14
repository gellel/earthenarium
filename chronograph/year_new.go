package chronograph

func NewYear(year int) *Year {
	return &Year{
		Leap:   (year%4 == 0 && year%100 != 0 || year%400 == 0),
		Number: year}
}
