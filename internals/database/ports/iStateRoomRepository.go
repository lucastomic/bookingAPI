package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IStateRoomRepository interface {
	repository[domain.StateRoom, int]
	FindByBoatId(int) ([]domain.StateRoom, error)
}
