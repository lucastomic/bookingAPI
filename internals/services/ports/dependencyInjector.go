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
	repo := databaseport.NewStateRoomRepository()
	return services.StateRoomService{IStateRoomRepository: repo}
}

func NewReservationService() IReservationService {
	repo := databaseport.NewReservationRepository()
	return services.ReservationServcie{IReservationRepository: repo}
}
