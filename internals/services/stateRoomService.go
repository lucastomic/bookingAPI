package services

import (
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// StateRoomService is a service that provides operations related to state rooms.
type StateRoomService struct {
	databaseport.IStateRoomRepository
}

// AddStateRoom adds a new state room by calling the UpdateStateRoom() method with the given state room,
// and returns the updated state room or an error if the update fails.
func (s StateRoomService) AddStateRoom(stateRoom domain.StateRoom) (domain.StateRoom, error) {
	return s.UpdateStateRoom(stateRoom)
}

// UpdateStateRoom updates an existing state room by calling the Save() method with the given state room,
// and returns the updated state room or an error if the save operation fails.
func (s StateRoomService) UpdateStateRoom(stateRoom domain.StateRoom) (domain.StateRoom, error) {
	err := s.Save(stateRoom)
	if err != nil {
		return *domain.EmptyStateRoom(), err
	}
	return stateRoom, nil
}

// DeleteStateRoom deletes a state room by calling the Remove() method with the given state room,
// and returns an error if the removal operation fails.
func (s StateRoomService) DeleteStateRoom(stateRoom domain.StateRoom) error {
	err := s.Remove(stateRoom)
	if err != nil {
		return err
	}
	return nil
}
