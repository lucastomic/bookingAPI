package domain

import (
	"strconv"

	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

// Reservation struct represents a booking reservation for a boat state room.
type Reservation struct {
	id          int
	maxCapacity int
	firstDay    timesimplified.Time
	lastDay     timesimplified.Time
	clients     []*Client
	isOpen      bool
	boatId      int
}

func (r Reservation) CanMerge(reservation Reservation) bool {
	bothReservationsAreOpen := r.bothAreOpen(reservation)
	datesMatch := r.datesMatch(reservation)
	dontExceedsMaximumCapacityWith := !r.exceedsMaximumCapacityWith(reservation.clients...)
	return bothReservationsAreOpen && datesMatch && dontExceedsMaximumCapacityWith
}

func (r Reservation) CanMergePassengers(passengers int) bool {
	dontExceedsMaximumCapacityWith := !r.exceedsMaximumCapacityWith(&Client{0, "", "", passengers})
	return r.isOpen && dontExceedsMaximumCapacityWith
}

func (r *Reservation) Merge(reservation Reservation) {
	if !r.CanMerge(reservation) {
		return
	}
	r.addClients(reservation.clients...)
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
	return (r.firstDay.Before(dateToCheck) && r.lastDay.After(dateToCheck)) ||
		dateToCheck.Equals(r.firstDay) ||
		dateToCheck.Equals(r.lastDay)
}

// Overlaps checks whether the reservation r collides with the given reservation reservationToCheck. This means, this reservations
// shares at least one day
func (r Reservation) Overlaps(reservationToCheck Reservation) bool {
	return r.Contains(reservationToCheck.firstDay) || r.Contains(reservationToCheck.lastDay) ||
		reservationToCheck.Contains(r.firstDay) ||
		reservationToCheck.Contains(r.lastDay)
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
	return reservation.id == r.id && reservation.firstDay.Equals(r.firstDay) &&
		reservation.lastDay.Equals(r.lastDay) &&
		reservation.maxCapacity == r.maxCapacity &&
		r.isOpen == reservation.isOpen &&
		r.boatId == reservation.boatId &&
		r.getTotalPassengers() == reservation.getTotalPassengers()
}

func (r Reservation) getTotalPassengers() int {
	response := 0
	for _, client := range r.clients {
		response += client.passengers
	}
	return response
}

func (r Reservation) datesMatch(reservation Reservation) bool {
	return reservation.firstDay.Equals(r.firstDay) && reservation.lastDay.Equals(r.lastDay)
}

func (r Reservation) exceedsMaximumCapacityWith(clients ...*Client) bool {
	totalPassengers := r.getTotalPassengers()
	for _, client := range clients {
		totalPassengers += client.passengers
	}
	return totalPassengers > r.maxCapacity
}

func (r Reservation) bothAreOpen(reservation Reservation) bool {
	return reservation.isOpen && r.isOpen
}

func (r *Reservation) addClients(clients ...*Client) {
	r.clients = append(r.clients, clients...)
}

// EmptyReservation returns a new empty Reservation struct pointer.
func EmptyReservation() *Reservation {
	return &Reservation{}
}

func (r Reservation) Clients() []*Client {
	return r.clients
}

func (r *Reservation) SetClients(clients []*Client) {
	r.clients = clients
}

func (r Reservation) MaxCapacity() int {
	return r.maxCapacity
}

func (r *Reservation) SetMaxCapacity(maxCap int) {
	r.maxCapacity = maxCap
}

func (r Reservation) FirstDay() timesimplified.Time {
	return r.firstDay
}

func (r Reservation) BoatId() int {
	return r.boatId
}

func (r Reservation) LastDay() timesimplified.Time {
	return r.lastDay
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

// NewReservation creates and returns a new Reservation struct pointer with the provided parameters.
func NewReservation(
	id int,
	firstDay timesimplified.Time,
	lastDay timesimplified.Time,
	client *Client,
	isOpen bool,
	boatId int,
) *Reservation {
	return &Reservation{
		id:       id,
		clients:  []*Client{client},
		firstDay: firstDay,
		lastDay:  lastDay,
		isOpen:   isOpen,
		boatId:   boatId,
	}
}

func NewReservationWithCapacity(
	id int,
	firstDay timesimplified.Time,
	lastDay timesimplified.Time,
	client *Client,
	isOpen bool,
	maxCapacity int,
	boatId int,
) *Reservation {
	return &Reservation{
		id:          id,
		maxCapacity: maxCapacity,
		clients:     []*Client{client},
		firstDay:    firstDay,
		lastDay:     lastDay,
		isOpen:      isOpen,
		boatId:      boatId,
	}
}

func NewReservationWithoutClient(
	id int,
	firstDay timesimplified.Time,
	lastDay timesimplified.Time,
	isOpen bool,
	boatId int,
) *Reservation {
	return &Reservation{
		id:       id,
		clients:  []*Client{},
		firstDay: firstDay,
		lastDay:  lastDay,
		boatId:   boatId,
		isOpen:   isOpen,
	}
}
