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

// AddEmptyStateRoom adds an empty sateRoom to a boat
func (s StateRoomService) AddEmptyStateRoom(boatId int) error {
	stateRoom := domain.EmptyStateRoom()
	stateRoom.SetBoatId(boatId)
	err := s.setStateRoomId(stateRoom)
	if err != nil {
		return err
	}
	_, err = s.UpdateStateRoom(*stateRoom)
	return err
}

// setStateRoomId sets the Id of the given stateRoom to the amount of staterooms
// that are in the stateRoom's boat (besides the stateroom treated). So, if the stateroom belongs to e boat
// with 3 staterooms + the stateRoom passed as argument, the ID of this stateRoom will be set to 3
func (s StateRoomService) setStateRoomId(stateRoom *domain.StateRoom) error {
	boatStateRooms, err := s.FindByBoatId(stateRoom.BoatId())
	if err != nil {
		return err
	}
	(*stateRoom).SetId(len(boatStateRooms))
	return nil
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
	return err
}
