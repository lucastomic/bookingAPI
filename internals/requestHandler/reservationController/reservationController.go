package reservationcontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	serviceinjector "github.com/lucastomic/naturalYSalvajeRent/internals/services/injection"
)

const reservationEndpoint = "reservation"
const removeReservationEndpoint = reservationEndpoint + "/:id"

var reservationService = serviceinjector.NewReservationService()

func AddEndpoints(r *gin.IRoutes) {
	(*r).DELETE(removeReservationEndpoint, removeReservation)
}

// removeReservation function removes a reservation by its id from the database through the reservationService.
// If any error occurs during the process, it returns an error response to the client.
func removeReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		exceptionhandling.HandleException(c, exceptions.WrongIdType)
	}
	reservation, err := reservationService.GetReservation(id)
	if err != nil {
		exceptionhandling.HandleException(c, err)
	}

	err = reservationService.DeleteReservation(reservation)
	if err != nil {
		exceptionhandling.HandleException(c, err)
	}
}
