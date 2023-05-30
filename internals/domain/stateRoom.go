package domain

import (
	"errors"

	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

// The StateRoom struct defines a state room in a boat.
// A StateRoom has an id, the id of the boat it belongs to, and a list of reservations made for it.
type StateRoom struct {
	id           int
	boatId       int
	reservations []Reservation
}

// Id returns the id of the state room.
func (s StateRoom) Id() int {
	return s.id
}

func (s *StateRoom) SetId(id int) {
	s.id = id
}

// BoatId returns the id of the boat the state room belongs to.
func (s StateRoom) BoatId() int {
	return s.boatId
}

// SetBoatId sets the id of the boat the state room belongs to.
func (s *StateRoom) SetBoatId(boatId int) {
	s.boatId = boatId
}

// Reservations returns the list of reservations made for the state room.
func (s StateRoom) Reservations() []Reservation {
	return s.reservations
}

// SetReservedDays sets the list of reservations made for the state room.
func (s *StateRoom) SetReservedDays(reservation []Reservation) {
	s.reservations = reservation
}

// AddReservation adds a new reservation to a stateroom. If the reservation collides with another
// reservation already reserved, it throws an error
func (s *StateRoom) AddReservation(reservation Reservation) error {
	if !s.ableToReservate(reservation) {
		return errors.New("reservation collides with another reservation")
	}
	reservation.SetStateRoomId(s.id)
	s.reservations = append(s.reservations, reservation)
	return nil
}

// GetStartedReservation returns the current reservation (those which has already started but hasn't finished yet)
// if it exists. If it doesn't existm returns a zero reservation
func (s StateRoom) GetStartedReservation() Reservation {
	for _, stateRoomReservation := range s.reservations {
		if stateRoomReservation.Contains(timesimplified.Now()) {
			return stateRoomReservation
		}
	}
	return *new(Reservation)
}

// RemoveReservation removes the given reservation. If the specified reservation is not
// in the stateRoom it throws an error
func (s *StateRoom) RemoveReservation(reservationToRemve Reservation) error {
	for i, reservation := range s.reservations {
		if reservation.Equals(reservationToRemve) {
			if len(s.reservations)-1 == i {
				s.reservations = s.reservations[:i]
			} else {
				s.reservations = append(s.reservations[:i], s.reservations[i+1:]...)
			}
			return nil
		}
	}
	return errors.New("reservation doesn't exist")
}

// ableToReservate checks whether the range of a given reservation is completly free for reservate in the
// stateroom
func (s *StateRoom) ableToReservate(reservationToCheck Reservation) bool {
	for _, stateRoomReservation := range s.reservations {
		if stateRoomReservation.Overlaps(reservationToCheck) {
			return false
		}
	}
	return true

}

// EmptyStateRoom creates and returns a new empty StateRoom.
func EmptyStateRoom() *StateRoom {
	return &StateRoom{}
}

// NewStateRoom creates and returns a new StateRoom with the given parameters.
func NewStateRoom(id int, boatId int, reservedDays []Reservation) *StateRoom {
	return &StateRoom{id, boatId, reservedDays}
}
