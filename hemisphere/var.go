package hemisphere

var (
	Antarctic = NewArea(south, ((south + antarctic) / 2), antarctic)
)

var (
	Arctic = NewArea(north, ((north + arctic) / 2), arctic)
)

var (
	Cancer = NewArea(arctic, ((cancer + arctic) / 2), cancer)
)

var (
	Capricorn = NewArea(antarctic, ((capricorn + antarctic) / 2), capricorn)
)

var (
	East = NewArea(east, (east / 2), prime)
)

var (
	Equator = NewArea(cancer, equator, capricorn)
)

var (
	Prime = NewArea(east/2, prime, west/2)
)

var (
	North = NewArea(north, (north / 2), equator)
)

var (
	South = NewArea(south, (south / 2), equator)
)

var (
	West = NewArea(west, (west / 2), prime)
)

var Center = &Zone{
	equatorLabel: Equator,
	primeLabel:   Prime}

var Eastern = &Zone{
	eastLabel: East}

var Northern = &Zone{
	arcticLabel: Arctic,
	cancerLabel: Cancer,
	northLabel:  North}

var Southern = &Zone{
	antarcticLabel: Antarctic,
	capricornLabel: Capricorn,
	southLabel:     South}

var Western = &Zone{
	westLabel: West}

var Zones = (&Zone{}).Concat(
	Center,
	Eastern,
	Northern,
	Southern,
	Western)
