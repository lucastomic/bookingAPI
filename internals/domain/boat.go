package domain

import (
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
	ReservationHandler
}

func NewBoat(name string, stateRooms []*StateRoom, owner string, maxCapacity int) *Boat {
	return &Boat{name: name, stateRooms: stateRooms, owner: owner, maxCapacity: maxCapacity,
		ReservationHandler: ReservationHandler{stateRooms, maxCapacity},
	}
}

func NewBoatWithId(
	id int,
	name string,
	stateRooms []*StateRoom,
	owner string,
	maxCapacity int,
) *Boat {
	return &Boat{
		id:                 id,
		name:               name,
		stateRooms:         stateRooms,
		owner:              owner,
		maxCapacity:        maxCapacity,
		ReservationHandler: ReservationHandler{stateRooms, maxCapacity},
	}
}

func EmptyBoat() *Boat {
	return &Boat{}
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

func (b Boat) GetNotAvailableDaysForSharedReservation(passengers int) []timesimplified.Time {
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

func (b Boat) GetFullCapacityDays() []timesimplified.Time {
	daysCounter := dayscounter.NewDaysCounter(len(b.stateRooms))
	b.forEachReservation(func(reservation *Reservation) {
		reservation.ForEachDay(func(date timesimplified.Time) {
			daysCounter.Add(date)
		})
	})
	return daysCounter.GetWhichArchivedObjetive()
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

func (b Boat) MaxCapacity() int {
	return b.maxCapacity
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
