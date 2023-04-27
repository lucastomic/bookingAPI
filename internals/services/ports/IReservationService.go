package serviceports

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IReservationService interface {
	CreateReservation(reservation domain.Reservation) error
	UpdateReservation(reservation domain.Reservation) error
	DeleteReservation(reservation domain.Reservation) error
}
