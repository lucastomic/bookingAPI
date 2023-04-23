package domain

type StateRoom struct {
	id           int
	boatId       int
	reservedDays []Reservation
}

func (s StateRoom) Id() int {
	return s.id
}
func (s StateRoom) BoatId() int {
	return s.boatId
}

func (s StateRoom) ReservedDays() []Reservation {
	return s.reservedDays
}
func (s *StateRoom) SetReservedDays(reservation []Reservation) {
	s.reservedDays = reservation
}

func EmptyStateRoom() *StateRoom {
	return &StateRoom{}
}

func NewStateRoom(id int, boatId int, reservedDays []Reservation) *StateRoom {
	return &StateRoom{id, boatId, reservedDays}
}
