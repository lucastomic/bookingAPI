package boatservicetests

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/services"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/testing/mockups"
)

var boatService = func() serviceports.IBoatService {
	repo := mockups.BoatMockUp{}
	reservationRepo := database.NewReservationRepository()
	return services.NewBoatService(repo, reservationRepo)
}()

var user1 = domain.NewClient("Lucas Tomic", "1234212", 0)
