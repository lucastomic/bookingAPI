package timesimplified

import (
	"time"
)

type Time time.Time

const dateFormat = "2006-01-02"

func NewTime(year int, month int, day int) Time {
	return Time(time.Date(year, time.Month(month), day, 20, 34, 58, 651387237, time.UTC))
}

// ToString returns a string representing a given Time. It only returns the
// yyyy-mm-dd format, this means, only the first 10 characters of the time.Time.String() method.
// "2023-05-03( 11:07:37.88739 +0200 CEST m=+0.001215626)"
func (t Time) ToString() string {
	return time.Time(t).String()[:10]
}

func (t Time) Day() int {
	return time.Time(t).Day()
}
func (t Time) Month() time.Month {
	return time.Time(t).Month()
}
func (t Time) Year() int {
	return time.Time(t).Year()
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t Time) AddDays(days int) Time {
	daysNumber := time.Duration(days)
	return Time(time.Time(t).Add(daysNumber * 24 * time.Hour))
}

func (t Time) Before(dateToCheck Time) bool {
	dateToCheckParsed := time.Time(dateToCheck)
	tParsed := time.Time(t)
	return tParsed.Before(dateToCheckParsed)
}
func (t Time) After(dateToCheck Time) bool {
	dateToCheckParsed := time.Time(dateToCheck)
	tParsed := time.Time(t)
	return tParsed.After(dateToCheckParsed)
}

func (t Time) Unix() int64 {
	tParsed := time.Time(t)
	return tParsed.Unix()
}

func Now() Time {
	return Time(time.Now())
}

func (t Time) Equals(other interface{}) bool {
	if otherTime, ok := other.(Time); ok {
		return t.Year() == otherTime.Year() && t.Month() == otherTime.Month() && t.Day() == otherTime.Day()
	}
	return false
}

// FromString returns a new Time given a string with the yyyy-mm-dd format
func FromString(in string) (Time, error) {
	timeParsed, err := time.Parse(dateFormat, in)
	if err != nil {
		return Time(time.Time{}), err
	}
	return Time(timeParsed), nil

}
