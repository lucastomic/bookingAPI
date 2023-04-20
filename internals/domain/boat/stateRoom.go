package boat

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain/reservation"

type StateRoom struct {
	id           int
	reservedDays []reservation.Reservation
}

func (s StateRoom) Id() int {
	return s.id
}

func (s StateRoom) ReservedDays() []reservation.Reservation {
	return s.reservedDays
}
