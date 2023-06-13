package boatservicetests

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	serviceinjector "github.com/lucastomic/naturalYSalvajeRent/internals/services/injection"
)

// TODO MUST BE REPLACED BY A BOATSERVICE WITH A MOCKED REPOSITORY
var boatService = serviceinjector.NewBoatService()
var user1 = domain.NewClient("Lucas Tomic", "1234212")
