package serviceports

import (
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/services"
)

func NewBoatService() IBoatService {
	repo := databaseport.NewBoatRepository()
	return services.NewBoatService(repo)
}

func NewStateRoomService() IStateRoomService {
	return services.StateRoomService{}
}

func NewReservationService() IReservationService {
	return services.ReservationServcie{}
}
