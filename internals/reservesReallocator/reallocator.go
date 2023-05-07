package reservesreallocator

import (
	"errors"

	"github.com/lucastomic/naturalYSalvajeRent/internals/datastructure"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

func RealloacteReserves(boat *domain.Boat, reservation *domain.Reservation) error {
	var success bool
	reservations := append(boat.GetUnstartedReservations(), *reservation)
	var reservationsQueue = datastructure.NewQueue(reservations)
	recursiveRealloaction(&success, boat, reservationsQueue)
	if !success {
		return errors.New("unable to reallocate new reservation")
	}
	return nil
}

func recursiveRealloaction(
	success *bool,
	boat *domain.Boat,
	reservations *datastructure.Queue[domain.Reservation],
) {
	if reservations.IsEmpty() {
		*success = true
	} else {
		i := 0
		stateRooms := boat.StateRooms()
		reservation, _ := reservations.Pop()
		for !*success && len(stateRooms) > i {
			if err := stateRooms[i].AddReservation(reservation); err == nil {
				oldBoatId := reservation.BoatId()
				oldStateRoomId := reservation.StateRoomId()
				reservation.SetBoatId(stateRooms[i].BoatId())
				reservation.SetStateRoomId(stateRooms[i].Id())
				recursiveRealloaction(success, boat, reservations)
				if !*success {
					reservation.SetBoatId(oldBoatId)
					reservation.SetStateRoomId(oldStateRoomId)
					stateRooms[i].RemoveReservation(reservation)
					reservations.Push(reservation)
				}
			}

		}
	}
}
