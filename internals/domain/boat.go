package domain

import (
	"errors"

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

// This method returns a slice of all unstarted reservations on the boat.
func (b Boat) GetUnstartedReservations() []*Reservation {
	var response []*Reservation
	for _, stateRoom := range b.stateRooms {
		for i, reservation := range stateRoom.Reservations() {
			if reservation.StartsAfter(timesimplified.Now()) {
				response = append(response, stateRoom.Reservations()[i])
			}
		}
	}
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

// GetStateRoomsWithStartedReservations retrieves the boat's staterooms with only thje reservations which has already started
// It doesn't modifies the actual staterooms. Only returns a copy of them with the started reservations
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

// GetNotEmptyDays retruns those days when there is at least one resrevation
func (b Boat) GetNotEmptyDays() []timesimplified.Time {
	hashMap := make(map[timesimplified.Time]bool)
	var response []timesimplified.Time
	for _, stateroom := range b.stateRooms {
		for _, reservation := range stateroom.reservations {
			reservation.ForEachDay(func(t timesimplified.Time) {
				if alreadyCounted := hashMap[t]; !alreadyCounted {
					hashMap[t] = true
					response = append(response, t)
				}
			})
		}
	}
	return response
}

// GetFullCapacityDays get a slice of days when all the boat's staterooms are reserved
func (b Boat) GetFullCapacityDays() []timesimplified.Time {
	var response []timesimplified.Time
	var daysHash map[timesimplified.Time]int = make(map[timesimplified.Time]int)
	for _, stateRoom := range b.StateRooms() {
		for _, reservation := range stateRoom.Reservations() {
			reservation.ForEachDay(func(date timesimplified.Time) {
				b.updateHashDays(&daysHash, &response, date)
			})
		}
	}
	return response
}

// updateHashDays takes a date and inserts it in the given hash map. If it already exists, it increments its position,
// if it doesn't is inserted with a value of 1. If any date get the same value as the amount of staterooms in the given boat,
// it inserts this date as a string in a string slice specified as parameter
func (b Boat) updateHashDays(
	daysHash *map[timesimplified.Time]int,
	response *[]timesimplified.Time,
	date timesimplified.Time,
) {
	if _, ok := (*daysHash)[date]; ok {
		(*daysHash)[date]++
		if (*daysHash)[date] == len(b.StateRooms()) {
			*response = append(*response, date)
		}
	} else {
		(*daysHash)[date] = 1
	}
}

// This method sets the state rooms of the boat.
func (b *Boat) SetStateRooms(stateRooms []*StateRoom) {
	for _, stateroom := range stateRooms {
		for _, reservation := range stateroom.reservations {
			reservation.maxCapacity = b.maxCapacity
		}
	}
	b.stateRooms = stateRooms
}
