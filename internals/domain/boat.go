package domain

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

type Boat struct {
	id         int
	name       string
	stateRooms []StateRoom
}

// This function creates a new Boat instance with the provided name and state rooms.
func NewBoat(name string, stateRooms []StateRoom) *Boat {
	return &Boat{name: name, stateRooms: stateRooms}
}

// This function creates a new Boat instance with the provided id, name and state rooms.
func NewBoatWithId(id int, name string, stateRooms []StateRoom) *Boat {
	return &Boat{id: id, name: name, stateRooms: stateRooms}
}

// This function creates a new empty Boat instance.
func EmtyBoat() *Boat {
	return &Boat{}
}

// This method returns the ID of the boat.
func (b Boat) Id() int {
	return b.id
}

// This method returns the name of the boat.
func (b Boat) Name() string {
	return b.name
}

// This method sets the name of the boat.
func (b *Boat) SetName(name string) {
	b.name = name
}

// This method returns a slice of all state rooms on the boat.
func (b Boat) StateRooms() []StateRoom {
	return b.stateRooms
}

// This method returns a slice of all unstarted reservations on the boat.
func (b Boat) GetUnstartedReservations() []*Reservation {
	var response []*Reservation
	for _, stateRoom := range b.stateRooms {
		for i, reservation := range stateRoom.Reservations() {
			if reservation.StartsAfter(timesimplified.Now()) {
				response = append(response, &stateRoom.Reservations()[i])
			}
		}
	}
	return response
}

// AddReservation looks for a free date's range in all the boat's stateRooms which matchs with the reservation
// one. If there is place to set the reservation, it adds to the stateroom and change the reservation's stateRoomId.
// It doesn't reallocates any reservation of the boat. In others words, does NOT change any reservation already reserved on the boat
// Returns true if the reservation was allocated propperly and false if there is no any free range for the reservation
func (b *Boat) AddReservation(reservation *Reservation) bool {
	couldReserve := false
	i := 0
	for i < len(b.StateRooms()) && !couldReserve {
		stateRoom := &b.StateRooms()[i]
		err := stateRoom.AddReservation(*reservation)
		couldReserve = err == nil
		i++
	}
	return couldReserve
}

// GetStateRoomsWithStartedReservations retrieves the boat's staterooms with only thje reservations which has already started
// It doesn't modifies the actual staterooms. Only returns a copy of them with the started reservations
func (b Boat) GetStateRoomsWithStartedReservations() []StateRoom {
	var response []StateRoom
	for _, stateRoom := range b.stateRooms {
		if reservation := stateRoom.GetStartedReservation(); !reservation.IsZero() {
			stateRoom.SetReservedDays([]Reservation{stateRoom.GetStartedReservation()})
		} else {
			stateRoom.SetReservedDays([]Reservation{})
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
func (b Boat) updateHashDays(daysHash *map[timesimplified.Time]int, response *[]timesimplified.Time, date timesimplified.Time) {
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
func (b *Boat) SetStateRooms(stateRooms []StateRoom) {
	b.stateRooms = stateRooms
}
