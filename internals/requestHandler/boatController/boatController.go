package boatController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

const boatEndpoint = "boat"
const getBoatEndpoint = boatEndpoint + "/:id"
const createBoatEndpoint = boatEndpoint

var boatService = serviceports.NewBoatService()

// UpdateEndpoints takes a gin.Engine object and updates all the boat endpoints
func AddEndpoints(r *gin.Engine) {
	r.GET(getBoatEndpoint, getBoat)
	r.POST(createBoatEndpoint, createBoat)
}

// createBoat creates a new boat given the
func createBoat(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&body); err != nil {
		badParameter(c, err)
		return
	}
	boat, err := boatService.CreateBoat(*domain.NewBoat(body.Name, []domain.StateRoom{}))
	if err != nil {
		unknownError(c, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"id":   boat.Id(),
		"name": boat.Name(),
	})

}

func badParameter(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error in parameter": err.Error()})
}

// getBoat retrieves a boat given its ID (by param) and updates the given gin.Context
// with the boat in JSON format
// In case of wrong ID by param (something different than a number) or other unkonwn error, it returns it as json.
func getBoat(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		invalidIdParameter(c)
		return
	}
	boat, err := boatService.GetBoat(idParsed)
	if err != nil {
		unknownError(c, err)
		return
	}
	parseBoat(c, boat)
}

// invalidIdParameter updates the given gin.context with a json explaning
// that a param has been passed with the wrong type
func invalidIdParameter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid id parameter type",
	})
}

// unknownError updates the given gin.context with a json explaning
// the given error
func unknownError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err,
	})
}

// parseBoat updates the given gin.context with the boat specified as parmeter
func parseBoat(c *gin.Context, boat domain.Boat) {
	if boat.Name() == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "boat not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":      boat.Name(),
		"id":        boat.Id(),
		"stateRoom": boat.StateRooms(),
	})
}
