package domain

// The StateRoom struct defines a state room in a boat.
// A StateRoom has an id, the id of the boat it belongs to, and a list of reservations made for it.
type StateRoom struct {
	id           int
	boatId       int
	reservedDays []Reservation
}

// Id returns the id of the state room.
func (s StateRoom) Id() int {
	return s.id
}

// BoatId returns the id of the boat the state room belongs to.
func (s StateRoom) BoatId() int {
	return s.boatId
}

// SetBoatId sets the id of the boat the state room belongs to.
func (s *StateRoom) SetBoatId(boatId int) {
	s.boatId = boatId
}

// ReservedDays returns the list of reservations made for the state room.
func (s StateRoom) ReservedDays() []Reservation {
	return s.reservedDays
}

// SetReservedDays sets the list of reservations made for the state room.
func (s *StateRoom) SetReservedDays(reservation []Reservation) {
	s.reservedDays = reservation
}

// EmptyStateRoom creates and returns a new empty StateRoom.
func EmptyStateRoom() *StateRoom {
	return &StateRoom{}
}

// NewStateRoom creates and returns a new StateRoom with the given parameters.
func NewStateRoom(id int, boatId int, reservedDays []Reservation) *StateRoom {
	return &StateRoom{id, boatId, reservedDays}
}
