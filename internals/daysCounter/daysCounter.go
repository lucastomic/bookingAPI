package dayscounter

import "github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"

type DaysCounter struct {
	objective          int
	archivedObjetive   []timesimplified.Time
	daysToReservations map[timesimplified.Time]int
}

func NewDaysCounter(objective int) *DaysCounter {
	return &DaysCounter{
		objective:          objective,
		archivedObjetive:   []timesimplified.Time{},
		daysToReservations: make(map[timesimplified.Time]int),
	}
}

func (b *DaysCounter) Add(
	date timesimplified.Time,
) {
	b.daysToReservations[date]++
	if b.daysToReservations[date] == b.objective {
		b.archivedObjetive = append(b.archivedObjetive, date)
	}
}

func (d DaysCounter) GetWhichArchivedObjetive() []timesimplified.Time {
	return d.archivedObjetive
}
