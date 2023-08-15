package database

import (
	boatDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/boat"
	clientDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/client"
	clientreservation "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/clientReservation"
	reservationDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/reservation"
	stateRoomDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/stateRoom"
	stateroomreservation "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/stateRoomReservation"
	userDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/user"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
)

func NewBoatRepository() databaseport.IBoatRepository {
	stateRoomRepo := NewStateRoomRepository()
	return boatDB.NewBoatRepository(stateRoomRepo)
}

func NewReservationRepository() databaseport.IReservationRepository {
	clientRepo := NewClientRepository()
	clientReservationRepo := clientreservation.Repository{}
	return reservationDB.NewReservationRepository(clientRepo, clientReservationRepo)
}

func NewStateRoomRepository() databaseport.IStateRoomRepository {
	reservationRepo := NewReservationRepository()
	stateroomReservation := stateroomreservation.Repository{}
	return stateRoomDB.NewStateRoomRepository(reservationRepo, stateroomReservation)
}

func NewUserRepository() databaseport.IUserRepository {
	boatRepo := NewBoatRepository()
	return userDB.NewUserRepository(boatRepo)
}

func NewClientRepository() databaseport.IClientRepository {
	return clientDB.NewClientRepository()
}
