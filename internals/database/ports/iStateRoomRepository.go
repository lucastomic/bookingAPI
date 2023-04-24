package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IStateRoomRepository interface {
	Save(domain.StateRoom) error
	FindById(...int) (domain.StateRoom, error)
	FindByBoatId(int) ([]domain.StateRoom, error)
}
