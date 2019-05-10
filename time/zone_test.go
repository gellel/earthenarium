package time_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gellel/earthenarium/time"
)

func TestNewZone(testing *testing.T) {
	const (
		e string = "time.NewZone(\"%s\") did not return \"%s\". returned \"%s\" instead."
	)
	for key := range time.Timezones {
		substrings := strings.Split(key, "")
		for i := 0; i < len(substrings); i++ {
			substring := substrings[i]
			if ok := (i%2 == 0); ok {
				substrings[i] = strings.ToUpper(substring)
			} else {
				substrings[i] = substring + "?@$@"
			}
		}
		namespace := strings.Join(substrings, "")
		zone := time.NewZone(namespace).String()
		if zone != key {
			testing.Fatal(fmt.Errorf(e, namespace, key, zone))
		}
	}
}
