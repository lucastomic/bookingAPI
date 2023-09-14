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
	b.forEachReservationDay(func(t timesimplified.Time, _ *Reservation) {
		days.AddIfNotExists(t)
	})
	return days.GetAsArray()
}

func (b DatesManager) GetNotAvailableDaysForSharedReservation(
	passengers int,
) []timesimplified.Time {
	response := timeset.NewTimeSet()
	b.forEachReservation(func(res *Reservation) {
		if !res.CanMergePassengers(passengers) {
			res.ForEachDay(func(day timesimplified.Time) {
				response.AddIfNotExists(day)
			})
		}
	})
	return response.GetAsArray()
}

func (b DatesManager) GetNotAvailableDaysForCloseReservation(
	stateroomsNeeded int,
) []timesimplified.Time {
	response := timeset.NewTimeSet()
	noEnoughStaterooms := b.getDaysCounterToNotEnoughStaterooms(stateroomsNeeded)
	b.forEachReservationDay(func(date timesimplified.Time, reservation *Reservation) {
		noEnoughStaterooms.Add(date)
		if reservation.isOpen {
			response.AddIfNotExists(date)
		}
	})
	response.AddSlice(noEnoughStaterooms.GetWhichArchivedObjetive())
	return response.GetAsArray()
}

func (b DatesManager) forEachReservation(function func(*Reservation)) {
	for _, stateRoom := range b.staterooms {
		for _, reservation := range stateRoom.Reservations() {
			function(reservation)
		}
	}
}

func (b DatesManager) forEachReservationDay(function func(timesimplified.Time, *Reservation)) {
	b.forEachReservation(func(res *Reservation) {
		res.ForEachDay(func(date timesimplified.Time) {
			function(date, res)
		})
	})
}

func (b DatesManager) getDaysCounterToNotEnoughStaterooms(
	stateroomsNeeded int,
) *dayscounter.DaysCounter {
	restantStaterooms := len(b.staterooms) - stateroomsNeeded
	notEnoughStaterooms := restantStaterooms + 1
	return dayscounter.NewDaysCounter(notEnoughStaterooms)
}
