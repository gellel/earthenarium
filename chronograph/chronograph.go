package chronograph

import (
	"fmt"
	"time"
)

type Day struct {
	Name   string
	Number int
	T      *time.Weekday
}

type Month struct {
	Name   string
	Number int
	T      *time.Month
}

type Span struct {
	Days    float64
	Hours   float64
	Minutes float64
	Months  float64
	Seconds float64
	T       *time.Duration
	Years   float64
}

type Time struct {
	Day       *Day
	Days      int
	Epoch     int64
	Month     *Month
	Second    int
	Time      *time.Time
	Timestamp string
	Year      int
	Zone      *Zone
}

type Zone struct {
	Name   string
	Offset int
}

func NewDay(timestamp *time.Time) *Day {
	t := timestamp.Weekday()
	return &Day{
		Name:   t.String(),
		Number: timestamp.Day(),
		T:      &t}
}

func NewMonth(timestamp *time.Time) *Month {
	t := timestamp.Month()
	return &Month{
		Name:   t.String(),
		Number: int(t),
		T:      &t}
}

func NewSpan(a *Time, b *Time) *Span {
	if a.Epoch < b.Epoch {
		return NewSpan(b, a)
	}
	duration := a.Time.Sub(*b.Time)
	seconds := duration.Seconds()
	minutes := duration.Minutes()
	hours := duration.Hours()
	days := (seconds / (60 * 60 * 24))
	years := (days / 365.25)
	return &Span{
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
		T:       &duration,
		Years:   years}
}

func NewTime(timestamp *time.Time) *Time {
	return &Time{
		Day:    NewDay(timestamp),
		Days:   timestamp.YearDay(),
		Epoch:  timestamp.Unix(),
		Month:  NewMonth(timestamp),
		Second: timestamp.Second(),
		Time:   timestamp,
		Year:   timestamp.Year(),
		Zone:   NewZone(timestamp)}
}

func NewTimeFromISO(str string) *Time {
	timestamp, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err)
	}
	return NewTime(&timestamp)
}

func NewZone(timestamp *time.Time) *Zone {
	name, offset := timestamp.Zone()
	return &Zone{
		Name:   name,
		Offset: offset}
}
