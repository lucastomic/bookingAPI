package boatservicetests

import (
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/services"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/testing/mockups"
)

// TODO MUST BE REPLACED BY A BOATSERVICE WITH A MOCKED REPOSITORY

var boatService = func() serviceports.IBoatService {
	repo := mockups.BoatMockUp{}
	reservationRepo := databaseport.NewReservationRepository()
	return services.NewBoatService(repo, reservationRepo)
}()

var user1 = domain.NewClient("Lucas Tomic", "1234212")
