package boatservicetests

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

var boatService = serviceports.NewBoatService()
var user1 = domain.NewUser("Lucas Tomic", "1234212")
