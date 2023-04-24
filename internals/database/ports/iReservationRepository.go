package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IReservationRepository interface {
	Save(domain.Reservation) error
	FindById(...int) (domain.Reservation, error)
	FindByStateRoom(domain.StateRoom) ([]domain.Reservation, error)
}
