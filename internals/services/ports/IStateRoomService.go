package serviceports

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IStateRoomService interface {
	AddStateRoom(stateRoom domain.StateRoom) (domain.StateRoom, error)
	UpdateStateRoom(stateRoom domain.StateRoom) (domain.StateRoom, error)
	DeleteStateRoom(stateRoom domain.StateRoom) error
	AddEmptyStateRoom(int) error
}
