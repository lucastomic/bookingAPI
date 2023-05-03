package domain

import (
	"time"

	"github.com/lucastomic/naturalYSalvajeRent/internals/timeParser"
)

// Reservation struct represents a booking reservation for a boat state room.
type Reservation struct {
	id          int
	user        User
	firstDay    time.Time
	lastDay     time.Time
	boatId      int
	stateRoomId int
}

// UserName returns the name of the user associated with the reservation.
func (r Reservation) UserName() string {
	return r.user.name
}

// UserPhone returns the phone number of the user associated with the reservation.
func (r Reservation) UserPhone() string {
	return r.user.phone
}

// FirstDay returns the start date of the reservation.
func (r Reservation) FirstDay() time.Time {
	return r.firstDay
}

// LastDay returns the end date of the reservation.
func (r Reservation) LastDay() time.Time {
	return r.lastDay
}

// BoatId returns the ID of the boat associated with the reservation.
func (r Reservation) BoatId() int {
	return r.boatId
}

// Id returns the unique ID of the reservation.
func (r Reservation) Id() int {
	return r.id
}

// StateRoomId returns the ID of the state room associated with the reservation.
func (r Reservation) StateRoomId() int {
	return r.stateRoomId
}

// EmptyReservation returns a new empty Reservation struct pointer.
func EmptyReservation() *Reservation {
	return &Reservation{}
}

// ForEachDay takes a function as a parameter and executes that function for each day of the reservation period.
func (r Reservation) ForEachDay(function func(time.Time)) {
	currentDate := r.FirstDay()
	for !timeParser.Equals(currentDate, r.LastDay()) {
		function(currentDate)
		currentDate = currentDate.Add(time.Hour * 24)
	}
}

// NewReservation creates and returns a new Reservation struct pointer with the provided parameters.
func NewReservation(id int, user User, firstDay time.Time, lastDay time.Time, boatId int, stateRoomId int) *Reservation {
	return &Reservation{
		id,
		user,
		firstDay,
		lastDay,
		boatId,
		stateRoomId,
	}
}
