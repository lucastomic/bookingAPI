package view

import (
	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type BoatView struct {
}

// ParseView takes a domain.Boat instance as parameters. It returns a map of type gin.H, which is a shorthand for the map[string]interface{} type in the Gin web framework.
// The function first initializes an empty slice of gin.H called parsedStateRooms and creates an instance of the stateRoomViewJSON struct.
// It then iterates over each state room in the boat's state rooms (accessed via the "StateRooms()" method of the "boat" parameter), and for each state room it calls
// the "ParseView" method of the "stateRoomView" variable with the current state room as the parameter.
// The resulting map is then appended to the "parsedStateRooms" slice.
// Finally, the function returns a map of type gin.H with the boat's name, ID, and the parsed state rooms.
func (bv BoatView) ParseView(boat domain.Boat) gin.H {
	var parsedStateRooms []gin.H
	stateRoomView := stateRoomViewJSON{}
	for _, stateRoom := range boat.StateRooms() {
		parsedStateRooms = append(parsedStateRooms, stateRoomView.ParseView(*stateRoom))
	}
	return gin.H{
		"name":      boat.Name(),
		"id":        boat.Id(),
		"owner":     boat.Owner(),
		"stateRoom": parsedStateRooms,
	}
}
