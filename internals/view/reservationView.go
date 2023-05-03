package view

import (
	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timeParser"
)

type ReservationViewJSON struct {
}

// ParseView takes in a ReservationViewJSON instance and a domain.Reservation instance as parameters.
// It returns a map of type gin.H, which is a shorthand for the map[string]interface{} type in the Gin web framework.
// The function creates a map using the "gin.H" type with the reservation's ID, user name, user phone, first day, and last day.
// These values are accessed using the respective methods on the "reservation" parameter.
// The resulting map is then returned by the function as the output. This function essentially maps the fields of a reservation domain
// object to a map object that can be returned as a JSON object using the Gin web framework.
func (view ReservationViewJSON) ParseView(reservation domain.Reservation) gin.H {
	return gin.H{
		"id":       reservation.Id(),
		"user":     reservation.UserName(),
		"phone":    reservation.UserPhone(),
		"firstDay": timeParser.ToString(reservation.FirstDay()),
		"lastDay":  timeParser.ToString(reservation.LastDay()),
	}
}
