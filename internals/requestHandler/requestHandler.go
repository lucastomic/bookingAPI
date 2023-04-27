package requesthandler

import (
	"github.com/gin-gonic/gin"

	"github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/boatController"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

var boatService = serviceports.NewBoatService()

func Run() {

	r := gin.Default()
	boatController.AddEndpoints(r)

	r.Run()
}
