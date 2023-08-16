package domain

import (
	"strconv"

	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

// Reservation struct represents a booking reservation for a boat state room.
type Reservation struct {
	id         int
	firstDay   timesimplified.Time
	lastDay    timesimplified.Time
	clients    []*Client
	passengers int
	isOpen     bool
	boatId     int
}

func (r Reservation) Clients() []*Client {
	return r.clients
}
func (r *Reservation) SetClients(clients []*Client) {
	r.clients = clients
}

// FirstDay returns the start date of the reservation.
func (r Reservation) FirstDay() timesimplified.Time {
	return r.firstDay
}

func (r Reservation) BoatId() int {
	return r.boatId
}

// LastDay returns the end date of the reservation.
func (r Reservation) LastDay() timesimplified.Time {
	return r.lastDay
}
func (r Reservation) Passengers() int {
	return r.passengers
}
func (r Reservation) IsOpen() bool {
	return r.isOpen
}

// Id returns the unique ID of the reservation.
func (r Reservation) Id() int {
	return r.id
}

func (r *Reservation) SetId(id int) {
	r.id = id
}

// IsZero checks whether the reservation is a zero value
func (s Reservation) IsZero() bool {
	return s.id == 0 && s.firstDay.IsZero() && s.lastDay.IsZero()

}

// String parses the reservation into a redeable string
func (s Reservation) String() string {
	var response string
	response += "clients:[" + "\n"
	for _, client := range s.clients {
		response += "	{" + "\n"
		response += "		name: " + client.Name() + "\n"
		response += "		phone: " + client.Phone() + "\n"
		response += "	}" + "\n"
	}
	response += "]" + "\n"
	response += "id: " + strconv.Itoa(s.Id()) + "\n"
	response += "first day: " + s.firstDay.ToString() + "\n"
	response += "last day: " + s.lastDay.ToString() + "\n"
	return response

}

// ForEachDay takes a function as a parameter and executes that function for each day of the reservation period.
func (r Reservation) ForEachDay(function func(timesimplified.Time)) {
	currentDate := r.FirstDay()
	for !currentDate.Equals(r.LastDay().AddDays(1)) {
		function(currentDate)
		currentDate = currentDate.AddDays(1)
	}
}

// Contains checks whether a concret day is contained in the reservation
func (r Reservation) Contains(dateToCheck timesimplified.Time) bool {
	return (r.firstDay.Before(dateToCheck) && r.lastDay.After(dateToCheck)) || dateToCheck.Equals(r.firstDay) || dateToCheck.Equals(r.lastDay)
}

// Overlaps checks whether the reservation r collides with the given reservation reservationToCheck. This means, this reservations
// shares at least one day
func (r Reservation) Overlaps(reservationToCheck Reservation) bool {
	return r.Contains(reservationToCheck.firstDay) || r.Contains(reservationToCheck.lastDay) || reservationToCheck.Contains(r.firstDay) || reservationToCheck.Contains(r.lastDay)
}

// HasStarted cheks whether the reservation has started yet (don't care if the reservation has already ended or not).
// Returns true if the first day of the reservation is before time.Now() and the first day is not today (time.Now())
func (r Reservation) HasStarted() bool {
	return r.firstDay.Before(timesimplified.Now()) && !r.firstDay.Equals(timesimplified.Now())
}

// EndsBefore checks whether r time range is finished before the given dateToCheck's day
func (r Reservation) EndsBefore(dateToCheck timesimplified.Time) bool {
	return r.lastDay.Before(dateToCheck) && (!dateToCheck.Equals(r.firstDay))
}

// StartsAfter checks whether r time range is started after the given dateToCheck's day
func (r Reservation) StartsAfter(dateToCheck timesimplified.Time) bool {
	return r.firstDay.After(dateToCheck) && (!dateToCheck.Equals(r.firstDay))
}

// Equals cheks whether the reservation is the same as the specified by argument.
func (r Reservation) Equals(reservation Reservation) bool {
	return reservation.id == r.id
}

// EmptyReservation returns a new empty Reservation struct pointer.
func EmptyReservation() *Reservation {
	return &Reservation{}
}

// NewReservation creates and returns a new Reservation struct pointer with the provided parameters.
func NewReservation(id int, firstDay timesimplified.Time, lastDay timesimplified.Time, client *Client, isOpen bool, passengers int, boatId int) *Reservation {
	return &Reservation{
		id:         id,
		clients:    []*Client{client},
		firstDay:   firstDay,
		lastDay:    lastDay,
		isOpen:     isOpen,
		passengers: passengers,
		boatId:     boatId,
	}
}
func NewReservationWithoutClient(id int, firstDay timesimplified.Time, lastDay timesimplified.Time, isOpen bool, passengers int, boatId int) *Reservation {
	return &Reservation{
		id:         id,
		clients:    []*Client{},
		firstDay:   firstDay,
		lastDay:    lastDay,
		isOpen:     isOpen,
		passengers: passengers,
		boatId:     boatId,
	}
}
