package domain

import (
	"time"
)

type Reservation struct {
	id          int
	user        User
	firstDay    time.Time
	lastDay     time.Time
	boatId      int
	stateRoomId int
}

func (r Reservation) UserName() string {
	return r.user.name
}
func (r Reservation) UserPhone() string {
	return r.user.phone
}
func (r Reservation) FirstDay() time.Time {
	return r.firstDay
}
func (r Reservation) LastDay() time.Time {
	return r.lastDay
}
func (r Reservation) BoatId() int {
	return r.boatId
}
func (r Reservation) Id() int {
	return r.id
}
func (r Reservation) StateRoomId() int {

	return r.stateRoomId
}

func EmptyReservation() *Reservation {
	return &Reservation{}
}

func NewReservation(id int, user User, firstDay time.Time, lastDay time.Time, boatId int, stateRoomId int) *Reservation {
	return &Reservation{
		id,
		user,
		firstDay,
		lastDay,
		boatId,
		stateRoomId,
	}
}
