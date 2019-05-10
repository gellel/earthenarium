package time

// Dayparts holds the supported arguments for use in time.NewDaypart.
var Dayparts = map[string]int{}

// Timezones holds the supported arguments for use in time.NewZone.
var Timezones = map[string]int{
	"Africa":     1,
	"America":    2,
	"Antarctica": 3,
	"Asia":       4,
	"Atlantic":   5,
	"Australia":  6,
	"Etc":        7,
	"Europe":     8,
	"Indian":     9,
	"Other":      10,
	"Pacific":    11}
