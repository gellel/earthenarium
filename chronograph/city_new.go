package chronograph

import (
	"regexp"
	"strings"
)

// NewCity instantiates a new City pointer.
func NewCity(city string) *City {
	reg := regexp.MustCompile("\\d")
	namespace := reg.ReplaceAllString(city, " ")
	reg = regexp.MustCompile("\\s{2,}")
	namespace = strings.Replace(namespace, "_", " ", -1)
	substrings := strings.Split(namespace, "/")
	for i := 0; i < len(substrings); i++ {
		substring := reg.ReplaceAllString(substrings[i], " ")
		substring = strings.TrimSpace(substring)
		subs := strings.Split(substring, " ")
		for j := 0; j < len(subs); j++ {
			sub := subs[j]
			if ok := len(sub) > 3; ok {
				sub = strings.ToLower(sub)
				subs[j] = strings.Title(sub)
			}
		}
		substrings[i] = strings.Join(subs, "_")
	}
	city = strings.Join(substrings, "/")
	c := City(city)
	return &c
}
