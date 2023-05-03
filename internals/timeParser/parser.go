package timeParser

import "time"

const dateFormat = "2006-01-02"

// ParseFromString takes a string with a date and converts it into a time.Time
// object, following the yyyy-mm-dd format
func ParseFromString(in string) (time.Time, error) {
	return time.Parse(dateFormat, in)
}

// Equals checks whether two dates are the same. It compares if both days, months and years are the same
func Equals(t1 time.Time, t2 time.Time) bool {
	return t1.Day() == t2.Day() && t1.Month() == t2.Month() && t1.Year() == t2.Year()
}

// ToString returns a string representing a given time.Time. It only returns the
// yyyy-mm-dd format, this means, only the first 10 characters of the time.Time.String() method.
// "2023-05-03( 11:07:37.88739 +0200 CEST m=+0.001215626)"
func ToString(t time.Time) string {
	return t.String()[:10]
}
