package domain

import (
	"errors"

	dayscounter "github.com/lucastomic/naturalYSalvajeRent/internals/daysCounter"
	timeset "github.com/lucastomic/naturalYSalvajeRent/internals/timeSet"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

type Boat struct {
	maxCapacity int
	owner       string
	id          int
	name        string
	stateRooms  []*StateRoom
}

func NewBoat(name string, stateRooms []*StateRoom, owner string, maxCapacity int) *Boat {
	return &Boat{name: name, stateRooms: stateRooms, owner: owner, maxCapacity: maxCapacity}
}

func NewBoatWithId(id int, name string, stateRooms []*StateRoom, owner string, maxCapacity int) *Boat {
	return &Boat{id: id, name: name, stateRooms: stateRooms, owner: owner, maxCapacity: maxCapacity}
}

func EmptyBoat() *Boat {
	return &Boat{}
}

func (b Boat) MaxCapacity() int {
	return b.maxCapacity
}

func (b Boat) Id() int {
	return b.id
}

func (b *Boat) SetId(id int) {
	b.id = id
}

func (b Boat) Owner() string {
	return b.owner
}

func (b Boat) Name() string {
	return b.name
}

func (b *Boat) SetName(name string) {
	b.name = name
}

func (b Boat) StateRooms() []*StateRoom {
	return b.stateRooms
}

func (b *Boat) SetStateRooms(stateRooms []*StateRoom) {
	stateRoomsModified := b.getWithMaxCapacity(stateRooms)
	b.stateRooms = stateRoomsModified
}

func (b Boat) GetUnstartedReservations() []*Reservation {
	var response []*Reservation
	b.forEachReservation(func(reservation *Reservation) {
		if reservation.StartsAfter(timesimplified.Now()) {
			response = append(response, reservation)
		}
	})
	return response
}

func (b *Boat) HasDisponibilityFor(reservation Reservation, stateroomsNeeded int) bool {
	availableStaterooms := 0
	for _, stateroom := range b.StateRooms() {
		if stateroom.CanReservate(reservation) {
			availableStaterooms++
		}
	}
	return availableStaterooms >= stateroomsNeeded
}

func (b *Boat) HasDisponibilityForEntireBoat(reservation Reservation) bool {
	return b.HasDisponibilityFor(reservation, len(b.stateRooms))
}

func (b *Boat) ReservateStateroom(reservation *Reservation) error {
	if !b.HasDisponibilityFor(*reservation, 1) {
		return errors.New("there is not enough space for this reservation")
	}
	if !reservation.isOpen {
		return errors.New("only close reservations can reserve only one stateroom. Shared resrevationss must reservate all of them")
	}
	i := 0
	for _, stateRoom := range b.StateRooms() {
		if (*stateRoom).CanReservate(*reservation) {
			reservation.SetMaxCapacity(b.maxCapacity)
			stateRoom.Reservate(reservation)
			break
		}
		i++
	}
	return nil
}

func (b *Boat) ReservateEveryStateroom(reservation *Reservation) error {
	if !b.HasDisponibilityForEntireBoat(*reservation) {
		return errors.New("there is not enough space for this reservation")
	}
	for _, stateroom := range b.StateRooms() {
		err := stateroom.Reservate(reservation)
		if err != nil {
			return err
		}
		reservation.SetMaxCapacity(b.maxCapacity)
	}
	return nil
}

func (b Boat) GetStateRoomsWithStartedReservations() []*StateRoom {
	var response []*StateRoom
	for _, stateRoom := range b.stateRooms {
		if reservation := stateRoom.GetStartedReservation(); !reservation.IsZero() {
			stateRoom.SetReservedDays([]*Reservation{stateRoom.GetStartedReservation()})
		} else {
			stateRoom.SetReservedDays([]*Reservation{})
		}
		response = append(response, stateRoom)
	}
	return response
}

func (b Boat) GetNotEmptyDays() []timesimplified.Time {
	days := timeset.NewTimeSet()
	b.forEachReservation(func(reservation *Reservation) {
		reservation.ForEachDay(func(t timesimplified.Time) {
			days.AddIfNotExists(t)
		})
	})
	return days.GetAsArray()
}

func (b Boat) GetDaysWithCloseReservations() []timesimplified.Time {
	days := timeset.NewTimeSet()
	b.forEachReservation(func(reservation *Reservation) {
		if !reservation.IsOpen() {
			reservation.ForEachDay(func(t timesimplified.Time) {
				days.AddIfNotExists(t)
			})
		}
	})
	return days.GetAsArray()
}

func (b Boat) GetFullCapacityDays() []timesimplified.Time {
	daysCounter := dayscounter.NewDaysCounter(len(b.stateRooms))
	b.forEachReservation(func(reservation *Reservation) {
		reservation.ForEachDay(func(date timesimplified.Time) {
			daysCounter.Add(date)
		})
	})
	return daysCounter.GetWhichArchivedObjetive()
}

func (b Boat) getWithMaxCapacity(stateRooms []*StateRoom) []*StateRoom {
	stateroomsModified := stateRooms
	for _, stateroom := range stateroomsModified {
		for _, reservation := range stateroom.reservations {
			reservation.maxCapacity = b.maxCapacity
		}
	}
	return stateroomsModified
}

func (b Boat) forEachReservation(function func(*Reservation)) {
	for _, stateRoom := range b.stateRooms {
		for _, reservation := range stateRoom.Reservations() {
			function(reservation)
		}
	}
}
