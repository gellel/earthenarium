package chronograph

import (
	"fmt"
	"regexp"
	"strconv"
)

// Timestamp is intended to be received as YYYY-MM-DDTHH-MM-SS.NNNZ.
type Timestamp string

// Correct evaluates the Timestamp string and checks whether it contains the required timestamp components.
// Does not validate whether the contents of the numeric bounds match the limits defined by the year.
func (timestamp *Timestamp) Correct() (ok bool) {
	content := string(*timestamp)
	reg, err := regexp.Compile("^[0-9]{4,}-[0-9]{2}-\\d{2}T([01]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9].[0-9]{3}Z")
	if ok = (err == nil); !ok {
		panic(err)
	}
	ok = reg.MatchString(content)
	return ok
}

// Day captures the day substring from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Day() (day string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("^\\d{4}-\\d{2}\\-(\\d{2})")
	if ok := (err == nil); !ok {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); !ok {
		return day
	}
	return captures[1]
}

// Error generates an error message for malformed Timestamp pointers.
func (timestamp *Timestamp) Error() error {
	return fmt.Errorf(errorTimestamp, string(*timestamp))
}

// Hour captures the hour substring from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Hour() (hour string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("T(\\d{2})")
	if ok := (err == nil); !ok {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); !ok {
		return hour
	}
	hour = captures[1]
	return hour
}

// Minute captures the minute substring from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Minute() (minute string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("T\\d{2}:(\\d{2})")
	if ok := (err == nil); !ok {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); !ok {
		return minute
	}
	minute = captures[1]
	return minute
}

// Month captures the fragment substring from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Month() (month string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("^\\d{4}-(\\d{2})")
	if ok := (err == nil); ok != true {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); ok != true {
		return month
	}
	month = captures[1]
	return month
}

// Nanosecond captures the millisecond substring from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Nanosecond() (millisecond string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("\\d{3}Z")
	if ok := (err == nil); !ok {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); !ok {
		return millisecond
	}
	millisecond = captures[1]
	return millisecond
}

// Parse parses a formatted string and returns the time value it represents.
func (timestamp *Timestamp) Parse() (year, month, day, hour, minute, second, nanosecond int) {
	year, _ = strconv.Atoi(timestamp.Year())
	month, _ = strconv.Atoi(timestamp.Month())
	day, _ = strconv.Atoi(timestamp.Day())
	hour, _ = strconv.Atoi(timestamp.Hour())
	minute, _ = strconv.Atoi(timestamp.Minute())
	second, _ = strconv.Atoi(timestamp.Second())
	nanosecond, _ = strconv.Atoi(timestamp.Nanosecond())
	return year, month, day, hour, minute, second, nanosecond
}

// Second captures the second fragment from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Second() (second string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("T\\d{2}:\\d{2}:(\\d{2})")
	if ok := (err == nil); ok != true {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); ok != true {
		return second
	}
	second = captures[1]
	return second
}

// String returns the Timestamp pointer a string.
func (timestamp *Timestamp) String() (time string) {
	time = string(*timestamp)
	return time
}

// Year captures the year fragment from the Timestamp pointer and returns it as a string.
func (timestamp *Timestamp) Year() (year string) {
	content := string(*timestamp)
	reg, err := regexp.Compile("^(\\d{4,})")
	if ok := (err == nil); ok != true {
		panic(err)
	}
	captures := reg.FindStringSubmatch(content)
	if ok := (len(captures) == 2); ok != true {
		return year
	}
	year = captures[1]
	return year
}
