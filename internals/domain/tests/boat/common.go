package boattest

import "github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"

func compareTimeSlices(s1 []timesimplified.Time, s2 []timesimplified.Time) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !contains(s1, s2[i]) {
			return false
		}
	}
	return true
}

func contains(s []timesimplified.Time, el timesimplified.Time) bool {
	for _, sEl := range s {
		if sEl.Equals(el) {
			return true
		}
	}
	return false
}
