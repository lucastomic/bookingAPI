package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IReservationRepository interface {
	repository[domain.Reservation, int]
	FindByStateRoom(domain.StateRoom) ([]domain.Reservation, error)
}
