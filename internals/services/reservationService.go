package services

import (
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// ReservationServcie is a service that provides operations related to reservations.
type ReservationServcie struct {
	databaseport.IReservationRepository
}

// CreateReservation creates a new reservation by calling the UpdateReservation() method with the given reservation,
// and returns an error if the update fails.
func (s ReservationServcie) CreateReservation(reservation domain.Reservation) error {
	return s.UpdateReservation(reservation)
}

// UpdateReservation updates an existing reservation by calling the Save() method with the given reservation,
// and returns an error if the save operation fails.
func (s ReservationServcie) UpdateReservation(reservation domain.Reservation) error {
	err := s.Save(reservation)
	if err != nil {
		return err
	}
	return nil
}

// DeleteReservation deletes a reservation by calling the Remove() method with the given reservation,
// and returns an error if the removal operation fails.
func (s ReservationServcie) DeleteReservation(reservation domain.Reservation) error {
	err := s.Remove(reservation)
	if err != nil {
		return err
	}
	return nil
}
