package requesthandler

import (
	"os"

	"github.com/gin-gonic/gin"

	authenticationcontroller "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/authenticationController"
	"github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/boatController"
	reservationcontroller "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/reservationController"
	stateroomcontroller "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/stateRoomController"
)

var (
	serverCert = os.Getenv("SERVER_CERT_PATH")
	serverKey  = os.Getenv("SERVER_KEY_PATH")
)

func Run() {

	r := gin.Default()
	boatController.AddEndpoints(r)
	stateroomcontroller.AddEndpoints(r)
	reservationcontroller.AddEndpoints(r)
	authenticationcontroller.AddEndpoints(r)

	// r.RunTLS(":8080", serverCert, serverKey)
	r.Run()
}
