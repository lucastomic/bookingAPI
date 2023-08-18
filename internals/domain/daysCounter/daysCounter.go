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
	if _, ok := (b.daysToReservations)[date]; ok {
		b.daysToReservations[date]++
		if b.daysToReservations[date] == b.objective {
			b.archivedObjetive = append(b.archivedObjetive, date)
		}
	} else {
		b.daysToReservations[date] = 1
	}
}

func (d DaysCounter) GetWhichArchivedObjetive() []timesimplified.Time {
	return d.archivedObjetive
}
