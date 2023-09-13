package domain

import (
	dayscounter "github.com/lucastomic/naturalYSalvajeRent/internals/daysCounter"
	timeset "github.com/lucastomic/naturalYSalvajeRent/internals/timeSet"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

type DatesManager struct {
	staterooms []*StateRoom
}

func (b DatesManager) GetUnstartedReservations() []*Reservation {
	var response []*Reservation
	b.forEachReservation(func(reservation *Reservation) {
		if reservation.StartsAfter(timesimplified.Now()) {
			response = append(response, reservation)
		}
	})
	return response
}

func (b DatesManager) GetNotEmptyDays() []timesimplified.Time {
	days := timeset.NewTimeSet()
	b.forEachReservation(func(reservation *Reservation) {
		reservation.ForEachDay(func(t timesimplified.Time) {
			days.AddIfNotExists(t)
		})
	})
	return days.GetAsArray()
}

func (b DatesManager) GetNotAvailableDaysForSharedReservation(
	passengers int,
) []timesimplified.Time {
	response := timeset.NewTimeSet()
	b.forEachReservation(func(res *Reservation) {
		if !res.isOpen || res.exceedsMaximumCapacityWith(&Client{0, "", "", passengers}) {
			res.ForEachDay(func(day timesimplified.Time) {
				response.AddIfNotExists(day)
			})
		}
	})
	return response.GetAsArray()
}

func (b DatesManager) GetFullCapacityDays() []timesimplified.Time {
	daysCounter := dayscounter.NewDaysCounter(len(b.staterooms))
	b.forEachReservation(func(reservation *Reservation) {
		reservation.ForEachDay(func(date timesimplified.Time) {
			daysCounter.Add(date)
		})
	})
	return daysCounter.GetWhichArchivedObjetive()
}

func (b DatesManager) forEachReservation(function func(*Reservation)) {
	for _, stateRoom := range b.staterooms {
		for _, reservation := range stateRoom.Reservations() {
			function(reservation)
		}
	}
}
