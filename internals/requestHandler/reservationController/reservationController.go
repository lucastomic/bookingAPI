package reservationcontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	reservationrequest "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/reservationController/reservationRequest"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

const reservationEndpoint = "reservation"
const addReservationEndpoint = reservationEndpoint
const removeReservationEndpoint = reservationEndpoint + "/:id"

var reservationService = serviceports.NewReservationService()

func AddEndpoints(r *gin.Engine) {
	r.POST(addReservationEndpoint, addReservation)
	r.DELETE(removeReservationEndpoint, removeReservation)
}

// addReservation function creates a new reservation by parsing the request body and validating the data.
// Then, it creates a new reservation and stores it in the database through the reservationService.
// If any error occurs during the process, it returns an error response to the client.
func addReservation(c *gin.Context) {
	var body reservationrequest.ReservationRequest
	if err := c.Bind(&body); err != nil {
		// TODO: HANDLE ERROR
		return
	}

	reservation, err := reservationService.ParseReservationRequest(body)
	err = reservationService.CreateReservation(reservation)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}

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
