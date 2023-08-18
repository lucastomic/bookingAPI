package timeset

import "github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"

type TimeSet struct {
	daysAlreadyCounted map[int64]bool
	set                []timesimplified.Time
}

func NewTimeSet() *TimeSet {
	return &TimeSet{
		make(map[int64]bool),
		[]timesimplified.Time{},
	}
}

func (t *TimeSet) AddIfNotExists(day timesimplified.Time) {
	alreadyCounted := t.daysAlreadyCounted[day.Unix()]
	if !alreadyCounted {
		t.daysAlreadyCounted[day.Unix()] = true
		t.set = append(t.set, day)
	}
}

func (t TimeSet) GetAsArray() []timesimplified.Time {
	return t.set
}
