package boatController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"

	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

const boatEndpoint = "boat"
const getBoatEndpoint = boatEndpoint + "/:id"
const createBoatEndpoint = boatEndpoint
const deleteBoatEndpoint = boatEndpoint + "/:id"

var boatService = serviceports.NewBoatService()

// AddEndpoints takes a gin.Engine object and updates all the boat endpoints
func AddEndpoints(r *gin.Engine) {
	r.GET(getBoatEndpoint, getBoat)
	r.POST(createBoatEndpoint, createBoat)
	r.DELETE(deleteBoatEndpoint, deleteBoat)
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

	c.JSON(http.StatusOK, gin.H{
		"name":      boat.Name(),
		"id":        boat.Id(),
		"stateRoom": boat.StateRooms(),
	})
}
