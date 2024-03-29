package services

import (
	"net/http"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"

	reservationrequest "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/reservationController/reservationRequest"
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
	err := s.Save(&reservation)
	if err != nil {
		return err
	}
	return nil
}

// GetReservation retrieves a reservation given its ID
func (s ReservationServcie) GetReservation(id int) (domain.Reservation, error) {
	reservation, err := s.FindById(id)
	if err != nil {
		return *domain.EmptyReservation(), err
	}
	if reservation.IsZero() {
		return *domain.EmptyReservation(), exceptions.NotFound
	}
	return reservation, nil

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

// ParseReservationRequest retrieeves a Reservation given a reservationRequest. If there is an error, it specifies
func (s ReservationServcie) ParseReservationRequest(req reservationrequest.ReservationRequest) (domain.Reservation, error) {
	firstDay, err := timesimplified.FromString(req.FirstDay)
	if err != nil {
		ex := exceptions.NewApiError(http.StatusBadRequest, "Bad firstDay format. Must be a string with yyyy-mm-dd format")
		return *domain.EmptyReservation(), ex
	}
	lastDay, err := timesimplified.FromString(req.LastDay)
	if err != nil {
		ex := exceptions.NewApiError(http.StatusBadRequest, "Bad lastDay format. Must be a string with yyyy-mm-dd format")
		return *domain.EmptyReservation(), ex
	}
	client := domain.NewClient(req.Email, req.Phone, req.Passengers)
	reservation := domain.NewReservation(0, firstDay, lastDay, client, req.IsOpen, req.BoatId)
	return *reservation, nil
}
