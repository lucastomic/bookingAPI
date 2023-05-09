package domain

import (
	"strconv"
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
func (r *Reservation) SetBoatId(id int) {
	r.boatId = id
}

// Id returns the unique ID of the reservation.
func (r Reservation) Id() int {
	return r.id
}

// StateRoomId returns the ID of the state room associated with the reservation.
func (r Reservation) StateRoomId() int {
	return r.stateRoomId
}

func (r *Reservation) SetStateRoomId(id int) {
	r.stateRoomId = id
}

// IsZero checks whether the reservation is a zero value
func (s Reservation) IsZero() bool {
	return s.id == 0 && s.boatId == 0 && s.user.name == "" && s.user.phone == "" && s.firstDay.IsZero() && s.stateRoomId == 0 && s.lastDay.IsZero()

}

// String parses the reservation into a redeable string
func (s Reservation) String() string {
	var response string
	response += "user name: " + s.UserName() + "\n"
	response += "user phone: " + s.UserPhone() + "\n"
	response += "boat: " + strconv.Itoa(s.BoatId()) + "\n"
	response += "id: " + strconv.Itoa(s.Id()) + "\n"
	response += "first day: " + timeParser.ToString(s.firstDay) + "\n"
	response += "last day: " + timeParser.ToString(s.lastDay) + "\n"
	return response

}

// ForEachDay takes a function as a parameter and executes that function for each day of the reservation period.
func (r Reservation) ForEachDay(function func(time.Time)) {
	currentDate := r.FirstDay()
	for !timeParser.Equals(currentDate, r.LastDay()) {
		function(currentDate)
		currentDate = currentDate.Add(time.Hour * 24)
	}
}

// Contains checks whether a concret day is contained in the reservation
func (r Reservation) Contains(dateToCheck time.Time) bool {
	return (r.firstDay.Before(dateToCheck) && r.lastDay.After(dateToCheck)) || timeParser.Equals(dateToCheck, r.firstDay) || timeParser.Equals(dateToCheck, r.lastDay)
}

// Overlaps checks whether the reservation r collides with the given reservation reservationToCheck. This means, this reservations
// shares at least one day
func (r Reservation) Overlaps(reservationToCheck Reservation) bool {
	return r.Contains(reservationToCheck.firstDay) || r.Contains(reservationToCheck.lastDay) || reservationToCheck.Contains(r.firstDay) || reservationToCheck.Contains(r.lastDay)
}

// EndsBefore checks whether r time range is finished before the given dateToCheck's day
func (r Reservation) EndsBefore(dateToCheck time.Time) bool {
	return r.lastDay.Before(dateToCheck) && (!timeParser.Equals(dateToCheck, r.firstDay))
}

// StartsAfter checks whether r time range is started after the given dateToCheck's day
func (r Reservation) StartsAfter(dateToCheck time.Time) bool {
	return r.firstDay.After(dateToCheck) && (!timeParser.Equals(dateToCheck, r.firstDay))
}

// Equals cheks whether the reservation is the same as the specified by argument.
func (r Reservation) Equals(reservation Reservation) bool {
	return reservation.id == r.id && reservation.boatId == r.boatId && r.firstDay == reservation.firstDay && r.user.name == reservation.user.name
}

// EmptyReservation returns a new empty Reservation struct pointer.
func EmptyReservation() *Reservation {
	return &Reservation{}
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
