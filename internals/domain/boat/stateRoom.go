package boat

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain/reservation"

type StateRoom struct {
	id           int
	boatId       int
	reservedDays []reservation.Reservation
}

func (s StateRoom) Id() int {
	return s.id
}
func (s StateRoom) BoatId() int {
	return s.boatId
}

func (s StateRoom) ReservedDays() []reservation.Reservation {
	return s.reservedDays
}

func EmptyStateRoom() *StateRoom {
	return &StateRoom{}
}

func NewStateRoom(id int, boatId int, reservedDays []reservation.Reservation) *StateRoom {
	return &StateRoom{id, boatId, reservedDays}
}
