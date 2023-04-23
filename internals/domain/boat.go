package domain

type Boat struct {
	id         int
	name       string
	stateRooms []StateRoom
}

func NewBoat(name string, stateRooms []StateRoom) *Boat {
	return &Boat{name: name, stateRooms: stateRooms}
}

func NewBoatWithId(id int, name string, stateRooms []StateRoom) *Boat {
	return &Boat{id: id, name: name, stateRooms: stateRooms}
}

func EmtyBoat() *Boat {
	return &Boat{}
}

func (b Boat) Id() int {
	return b.id
}

func (b Boat) Name() string {
	return b.name
}

func (b *Boat) SetName(name string) {
	b.name = name
}

func (b Boat) StateRooms() []StateRoom {
	return b.stateRooms
}

func (b *Boat) SetStateRooms(stateRooms []StateRoom) {
	b.stateRooms = stateRooms
}
