package stateroomcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	serviceinjector "github.com/lucastomic/naturalYSalvajeRent/internals/services/injection"
)

const stateRoomEndpoint = "stateRoom"
const addStateRoomEndpoint = stateRoomEndpoint + "/add/:boatId"

var stateRoomService = serviceinjector.NewStateRoomService()
var boatService = serviceinjector.NewBoatService()

// AddEndpoints takes a gin.Engine object and updates all the staetRoom endpoints
func AddEndpoints(r *gin.Engine) {
	r.PUT(addStateRoomEndpoint, addStateRoom)
}

// addStateRoom adds a state room to the boat whose ID is specified by parameter
// In case of success, doesn't return anything. Otherwise it return an error message
func addStateRoom(c *gin.Context) {
	boatId := c.Param("boatId")
	boatIdParsed, err := strconv.Atoi(boatId)
	if err != nil {
		err = exceptions.NewApiError(http.StatusBadRequest, "wrong boatId type. Mist be an integer")
		exceptionhandling.HandleException(c, err)
		return
	}
	boat, err := boatService.GetBoat(boatIdParsed)
	if err != nil || boat.Name() == "" {
		exceptionhandling.HandleException(c, err)
		return
	}
	err = stateRoomService.AddEmptyStateRoom(boatIdParsed)
	if err != nil {
		exceptionhandling.HandleException(c, err)
	}
}
