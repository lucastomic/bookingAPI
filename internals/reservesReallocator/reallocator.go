package reservesreallocator

import (
	"errors"

	"github.com/lucastomic/naturalYSalvajeRent/internals/datastructure"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// NOTE: The reallocating system is a Backtracking algorithm. To understand the way the reservations are reallocated,
// read: https://en.wikipedia.org/wiki/Backtracking

// ReallocateReserves realocate the given boat's reservations among the boat's stateroom in way that
// a new specified resrevation can be inserted in a stateroom.
// The reservations which has already started (those which contains today) can't be reallocated (we cant move a passenger
// who is alrady in the boat).
//
// If there is no way to reallocate the reservations to append the new one, it will throw an error
// explaining that reservation is impossble to be added
//
// For this reallocation it uses a Backtracking algorithm
func RealloacteReserves(boat *domain.Boat, reservation *domain.Reservation) error {
	var success bool
	reservations := append(boat.GetUnstartedReservations(), reservation)
	var reservationsQueue = datastructure.NewQueue(reservations)
	stateRooms := boat.GetStateRoomsWithStartedReservations()

	recursiveRealloaction(&success, &stateRooms, reservationsQueue)

	if !success {
		return errors.New("unable to set the new reservation. There is not enough space")
	} else {
		boat.SetStateRooms(stateRooms)
	}
	return nil
}

// recursiveRealloaction set sucesss to true if there is no more reservations to allocate. Otherwise,
// explore the different options to allocate the remaining reservations
func recursiveRealloaction(
	success *bool,
	stateRooms *[]*domain.StateRoom,
	reservations *datastructure.Queue[*domain.Reservation],
) {
	if reservations.IsEmpty() {
		*success = true
	} else {
		exploreChildNodes(success, stateRooms, reservations)
	}
}

// exploreChildNodes takes a reservation queue and explore the different options to allocate it in the given staterooms.
func exploreChildNodes(
	success *bool,
	stateRooms *[]*domain.StateRoom,
	reservations *datastructure.Queue[*domain.Reservation],
) {
	i := 0
	reservation, _ := reservations.Pop()
	for !*success && len(*stateRooms) > i {
		if err := (*stateRooms)[i].AddReservation(reservation); err == nil {
			recursiveRealloaction(success, stateRooms, reservations)
		}
		if !*success {
			(*stateRooms)[i].RemoveReservation(*reservation)
		}
		i++
	}
	if !*success {
		reservations.Push(reservation)
	}
}
