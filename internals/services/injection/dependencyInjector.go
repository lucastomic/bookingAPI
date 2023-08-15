package serviceinjector

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database"
	"github.com/lucastomic/naturalYSalvajeRent/internals/services"
	jwtservice "github.com/lucastomic/naturalYSalvajeRent/internals/services/jwtService"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

func NewBoatService() serviceports.IBoatService {
	repo := database.NewBoatRepository()
	reservationRepo := database.NewReservationRepository()
	return services.NewBoatService(repo, reservationRepo)
}

func NewStateRoomService() serviceports.IStateRoomService {
	repo := database.NewStateRoomRepository()
	return services.StateRoomService{IStateRoomRepository: repo}
}

func NewReservationService() serviceports.IReservationService {
	repo := database.NewReservationRepository()
	return services.ReservationServcie{IReservationRepository: repo}
}

func NewAuthenticationService() serviceports.AuthenticationService {
	repo := database.NewUserRepository()
	jwtService := NewJWTService()
	return services.NewAuthenticationService(repo, jwtService)
}

func NewJWTService() serviceports.JWTService {
	return jwtservice.NewJWTService()
}
