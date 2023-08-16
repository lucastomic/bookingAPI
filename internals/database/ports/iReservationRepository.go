package databaseport

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type IReservationRepository interface {
	Repository[domain.Reservation, int]
	FindByStateRoom(domain.StateRoom) ([]*domain.Reservation, error)
	FindByClient(client domain.Client) ([]*domain.Reservation, error)
}
