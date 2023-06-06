package boatservicetests

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

// TODO MUST BE REPLACED BY A BOATSERVICE WITH A MOCKED REPOSITORY
var boatService = serviceports.NewBoatService()
var user1 = domain.NewClient("Lucas Tomic", "1234212")
