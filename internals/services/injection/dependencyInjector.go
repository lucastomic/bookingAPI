package serviceinjector

import (
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/services"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

func NewBoatService() serviceports.IBoatService {
	repo := databaseport.NewBoatRepository()
	reservationRepo := databaseport.NewReservationRepository()
	return services.NewBoatService(repo, reservationRepo)
}

func NewStateRoomService() serviceports.IStateRoomService {
	repo := databaseport.NewStateRoomRepository()
	return services.StateRoomService{IStateRoomRepository: repo}
}

func NewReservationService() serviceports.IReservationService {
	repo := databaseport.NewReservationRepository()
	return services.ReservationServcie{IReservationRepository: repo}
}

func NewAuthenticationService() serviceports.AuthenticationService {
	repo := databaseport.NewUserRepository()
	jwtService := NewJWTService()
	return services.NewAuthenticationService(repo, jwtService)
}

func NewJWTService() serviceports.JWTService {
	return services.NewJWTService()
}
