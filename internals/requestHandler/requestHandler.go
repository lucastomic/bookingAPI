package requesthandler

import (
	"github.com/gin-gonic/gin"

	"github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/boatController"
	reservationcontroller "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/reservationController"
	stateroomcontroller "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/stateRoomController"
)

func Run() {

	r := gin.Default()
	boatController.AddEndpoints(r)
	stateroomcontroller.AddEndpoints(r)
	reservationcontroller.AddEndpoints(r)

	r.Run()
}
