package boat

type Boat struct {
	id         int
	name       string
	stateRooms []StateRoom
}

func NewBoat(id int, name string, stateRooms []StateRoom) *Boat {
	return &Boat{id, name, stateRooms}
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
func (b Boat) StateRooms() []StateRoom {
	return b.stateRooms
}
