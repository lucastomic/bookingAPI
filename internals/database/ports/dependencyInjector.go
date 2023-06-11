package databaseport

import (
	boatDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/boat"
	reservationDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/reservation"
	stateRoomDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/stateRoom"
	userDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/user"
)

func NewBoatRepository() BoatRepository {
	return boatDB.NewBoatRepository()
}

func NewReservationRepository() IReservationRepository {
	return reservationDB.NewReservationRepository()
}

func NewStateRoomRepository() IStateRoomRepository {
	return stateRoomDB.NewStateRoomRepository()
}
func NewUserRepository() UserRepository {
	return userDB.NewUserRepository()
}
