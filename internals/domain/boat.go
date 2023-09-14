package domain

type Boat struct {
	maxCapacity int
	owner       string
	id          int
	name        string
	stateRooms  []*StateRoom
	ReservationManager
	DatesManager
}

func NewBoat(name string, stateRooms []*StateRoom, owner string, maxCapacity int) *Boat {
	return &Boat{
		name: name, stateRooms: stateRooms,
		owner:              owner,
		maxCapacity:        maxCapacity,
		ReservationManager: ReservationManager{stateRooms, maxCapacity},
		DatesManager:       DatesManager{staterooms: stateRooms},
	}
}

func NewBoatWithId(
	id int,
	name string,
	stateRooms []*StateRoom,
	owner string,
	maxCapacity int,
) *Boat {
	return &Boat{
		id:                 id,
		name:               name,
		stateRooms:         stateRooms,
		owner:              owner,
		maxCapacity:        maxCapacity,
		ReservationManager: ReservationManager{stateRooms, maxCapacity},
		DatesManager:       DatesManager{staterooms: stateRooms},
	}
}

func EmptyBoat() *Boat {
	return &Boat{}
}

func (b Boat) GetStateRoomsWithStartedReservations() []*StateRoom {
	var response []*StateRoom
	for _, stateRoom := range b.stateRooms {
		if reservation := stateRoom.GetStartedReservation(); !reservation.IsZero() {
			stateRoom.SetReservedDays([]*Reservation{stateRoom.GetStartedReservation()})
		} else {
			stateRoom.SetReservedDays([]*Reservation{})
		}
		response = append(response, stateRoom)
	}
	return response
}

func (b Boat) Id() int {
	return b.id
}

func (b *Boat) SetId(id int) {
	b.id = id
}

func (b Boat) Owner() string {
	return b.owner
}

func (b Boat) Name() string {
	return b.name
}

func (b *Boat) SetName(name string) {
	b.name = name
}

func (b Boat) StateRooms() []*StateRoom {
	return b.stateRooms
}

func (b *Boat) SetStateRooms(stateRooms []*StateRoom) {
	stateRoomsModified := b.getWithMaxCapacity(stateRooms)
	b.stateRooms = stateRoomsModified
	b.DatesManager = DatesManager{staterooms: b.stateRooms}
	b.ReservationManager = ReservationManager{staterooms: b.stateRooms}
}

func (b Boat) MaxCapacity() int {
	return b.maxCapacity
}

func (b Boat) getWithMaxCapacity(stateRooms []*StateRoom) []*StateRoom {
	stateroomsModified := stateRooms
	for _, stateroom := range stateroomsModified {
		for _, reservation := range stateroom.reservations {
			reservation.maxCapacity = b.maxCapacity
		}
	}
	return stateroomsModified
}
