package view

import (
	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type stateRoomViewJSON struct {
}

// ParseView parse all the reservations into a JSON array and imbibe it into a stateRoom JSON
func (view stateRoomViewJSON) ParseView(stateRoom domain.StateRoom) gin.H {
	var parsedReservations []gin.H = []gin.H{}
	for _, reservation := range stateRoom.Reservations() {
		view.appendAsJson(&parsedReservations, reservation)
	}
	return gin.H{
		"id":           stateRoom.Id(),
		"reservedDays": parsedReservations,
	}
}

// appendAsJson converts the given reservation to JSON and appends it to the slice specified
func (view stateRoomViewJSON) appendAsJson(slice *[]gin.H, reservation domain.Reservation) {
	reservationView := ReservationViewJSON{}
	reservationJSON := reservationView.ParseView(reservation)
	*slice = append(*slice, reservationJSON)
}
