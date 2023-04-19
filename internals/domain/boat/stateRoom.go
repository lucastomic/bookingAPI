package boat

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain/reservation"

type StateRoom struct {
	reservedDays []reservation.Reservation
}
