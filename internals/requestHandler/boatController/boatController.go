package boatController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"

	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	reservationrequest "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/reservationController/reservationRequest"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
	viewport "github.com/lucastomic/naturalYSalvajeRent/internals/view/port"
)

const boatEndpoint = "boat"

const getBoatEndpoint = boatEndpoint + "/:id"
const createBoatEndpoint = boatEndpoint
const addReservationEndpoint = boatEndpoint + "/reservate"
const deleteBoatEndpoint = boatEndpoint + "/:id"
const getFullCapacityDaysEndpoint = boatEndpoint + "/reserved/:id"

var boatService = serviceports.NewBoatService()
var reservationService = serviceports.NewReservationService()
var boatView = viewport.NewBoatView()

// AddEndpoints takes a gin.Engine object and updates all the boat endpoints
func AddEndpoints(r *gin.Engine) {
	r.GET(getBoatEndpoint, getBoat)
	r.POST(createBoatEndpoint, createBoat)
	r.POST(addReservationEndpoint, addReservation)
	r.DELETE(deleteBoatEndpoint, deleteBoat)
	r.GET(getFullCapacityDaysEndpoint, getFullCapacityDays)
}

// createBoat receives a request to create a new boat, reads the boat name from the request body,
// creates a new boat object with the given name and an empty state room list,
// and then calls the boatService to save the new boat. If the creation is successful,
// it returns the created boat's id and name as a JSON response. Otherwise, it returns an error message.
func createBoat(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&body); err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}
	boat, err := boatService.CreateBoat(*domain.NewBoat(body.Name, []domain.StateRoom{}))
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"id":   boat.Id(),
		"name": boat.Name(),
	})
}

// deleteBoat receives a request to delete a boat with a specific id,
// retrieves the corresponding boat object from the boatService,
// and calls the boatService to delete it. If the deletion is successful,
// it returns a success message. Otherwise, it returns an error message.
func deleteBoat(c *gin.Context) {
	id := c.Param("id")

	idParsed, err := strconv.Atoi(id)
	if err != nil {
		exceptionhandling.HandleException(c, exceptions.WrongIdType)
		return
	}

	boat, err := boatService.GetBoat(idParsed)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}

	err = boatService.DeleteBoat(boat)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}
}

// getBoat receives a request to get the details of a boat with a specific id,
// retrieves the corresponding boat object from the boatService,
// and then calls parseBoat to return the boat details as a JSON response.
// If there is no boat with the given id or there is an error while retrieving the boat,
// it returns an error message.
func getBoat(c *gin.Context) {
	id := c.Param("id")

	idParsed, err := strconv.Atoi(id)
	if err != nil {
		exceptionhandling.HandleException(c, exceptions.WrongIdType)
		return
	}

	boat, err := boatService.GetBoat(idParsed)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}

	parseBoat(c, boat)
}

// parseBoat receives a boat object and returns the boat details,
// including the boat's id, name, and state rooms, as a JSON response.
// If the boat object is not found or there is an error while parsing the boat, it returns an error message.
func parseBoat(c *gin.Context, boat domain.Boat) {
	if boat.Name() == "" {
		exceptionhandling.HandleException(c, exceptions.NotFound)
		return
	}

	c.JSON(http.StatusOK, boatView.ParseView(boat))
}

// getFullCapacityDays retrieve the completed days of a boat given its ID.
// This means, the days when all the boat's staterooms are reserved
func getFullCapacityDays(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	id := c.Param("id")

	idParsed, err := strconv.Atoi(id)
	if err != nil {
		exceptionhandling.HandleException(c, exceptions.WrongIdType)
		return
	}

	boat, err := boatService.GetBoat(idParsed)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}

	completeDays := boatService.GetFullCapacityDays(boat)
	c.JSON(http.StatusOK, gin.H{
		"days": completeDays,
	})
}

// addReservation adds a new reservation to a boat.
// If there isn't enoguh space to reservate in the specified dates range, it returns an error.
// Also returns an error if the request body is not correct or if the boat id specified doesn't exist
func addReservation(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	reservation, err := parseReservationFromBody(c)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}

	if reservation.HasStarted() {
		err = exceptions.NewApiError(http.StatusBadRequest, "this dates are not allowed (first day has already passed)")
		exceptionhandling.HandleException(c, err)
		return
	}

	boat, err := boatService.GetBoat(reservation.BoatId())
	if err != nil {
		err = exceptions.NewApiError(http.StatusBadRequest, "boat with id "+strconv.Itoa(reservation.BoatId())+" doesn't exist")
		exceptionhandling.HandleException(c, err)
		return
	}
	err = boatService.AddReservation(boat, reservation)
	if err != nil {
		exceptionhandling.HandleException(c, err)
	}
}

// parseReservationFromBody parses the reservation from a request body. If there is any error, it returns it as
// second value with an empty reservation as first. Otherwise, returns the reservation
func parseReservationFromBody(c *gin.Context) (domain.Reservation, error) {
	var body reservationrequest.ReservationRequest
	if err := c.BindJSON(&body); err != nil {
		return *domain.EmptyReservation(), err
	}

	reservation, err := reservationService.ParseReservationRequest(body)
	if err != nil {
		return *domain.EmptyReservation(), err
	}
	return reservation, nil

}
