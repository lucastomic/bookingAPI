package databaseport

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type IStateRoomRepository interface {
	Repository[domain.StateRoom, int]
	FindByBoatId(int) ([]domain.StateRoom, error)
	FindByReservation(domain.Reservation) ([]domain.StateRoom, error)
}
