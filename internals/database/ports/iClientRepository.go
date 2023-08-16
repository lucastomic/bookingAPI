package databaseport

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type IClientRepository interface {
	Repository[domain.Client, int]
	FindByReservation(reservation domain.Reservation) ([]*domain.Client, error)
}
