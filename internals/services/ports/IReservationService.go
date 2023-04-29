package serviceports

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	reservationrequest "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/reservationController/reservationRequest"
)

type IReservationService interface {
	CreateReservation(reservation domain.Reservation) error
	UpdateReservation(reservation domain.Reservation) error
	DeleteReservation(reservation domain.Reservation) error
	GetReservation(id int) (domain.Reservation, error)
	ParseReservationRequest(reservationrequest.ReservationRequest) (domain.Reservation, error)
}
