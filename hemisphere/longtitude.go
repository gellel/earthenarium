package hemisphere

type Longtitude float32

func (longtitude Longtitude) East() bool {
	return longtitude < Eastern
}

func (longtitude Longtitude) Name() Label {
	if longtitude.East() {
		return EasternHemisphere
	}
	return WesternHemisphere
}

func (longtitude Longtitude) West() bool {
	return longtitude > Western
}
