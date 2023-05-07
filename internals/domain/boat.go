package domain

import "time"

type Boat struct {
	id         int
	name       string
	stateRooms []StateRoom
}

// This function creates a new Boat instance with the provided name and state rooms.
func NewBoat(name string, stateRooms []StateRoom) *Boat {
	return &Boat{name: name, stateRooms: stateRooms}
}

// This function creates a new Boat instance with the provided id, name and state rooms.
func NewBoatWithId(id int, name string, stateRooms []StateRoom) *Boat {
	return &Boat{id: id, name: name, stateRooms: stateRooms}
}

// This function creates a new empty Boat instance.
func EmtyBoat() *Boat {
	return &Boat{}
}

// This method returns the ID of the boat.
func (b Boat) Id() int {
	return b.id
}

// This method returns the name of the boat.
func (b Boat) Name() string {
	return b.name
}

// This method sets the name of the boat.
func (b *Boat) SetName(name string) {
	b.name = name
}

// This method returns a slice of all state rooms on the boat.
func (b Boat) StateRooms() []StateRoom {
	return b.stateRooms
}

// This method returns a slice of all unstarted reservations on the boat.
func (b Boat) GetUnstartedReservations() []Reservation {
	var response []Reservation
	for _, stateRoom := range b.stateRooms {
		for _, reservation := range stateRoom.Reservations() {
			if !reservation.Contains(time.Now()) {
				response = append(response, reservation)
			}
		}
	}
	return response
}

// This method sets the state rooms of the boat.
func (b *Boat) SetStateRooms(stateRooms []StateRoom) {
	b.stateRooms = stateRooms
}
