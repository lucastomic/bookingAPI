package timesimplified

import "time"

type Time time.Time

const dateFormat = "2006-01-02"

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

// Equals checks whether two dates are the same. It compares if both days,
// months and years are the same
func (t Time) Equals(timeToCompare Time) bool {
	return t.Day() == timeToCompare.Day() && t.Month() == timeToCompare.Month() && t.Year() == timeToCompare.Year()
}

// FromString returns a new Time given a string with the yyyy-mm-dd format
func FromString(in string) (Time, error) {
	timeParsed, err := time.Parse(dateFormat, in)
	if err != nil {
		return Time(time.Time{}), err
	}
	return Time(timeParsed), nil

}
