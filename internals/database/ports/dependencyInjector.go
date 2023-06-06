package databaseport

import (
	boatDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/boat"
	reservationDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/reservation"
	stateRoomDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/stateRoom"
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
