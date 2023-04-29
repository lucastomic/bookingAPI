package timeParser

import "time"

const dateFormat = "2006-01-02"

// ParseFromString takes a string with a date and converts it into a time.Time
// object, following the yyyy-mm-dd format
func ParseFromString(in string) (time.Time, error) {
	return time.Parse(in, dateFormat)
}
