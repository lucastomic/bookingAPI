# Booking API
Booking API is a simple API for reservation management in a boat with different staterooms. This is designed mainly for the porpouse of showing mu backend skills.
This is the stack used for the project development:
- Go
- Docker
- MYSQL
- AWS
Next, It will be explained the different software engeneering areas involved in this project
# Software engeneering areas involved
## Architecture
The architecture follows a MVC pattern, where the differents layers are connected through interfaces (dependency inversion) similar to a hexagonal architecture

![Diagram class](https://github.com/lucastomic/bookingAPI/assets/65186233/cdf91e60-1345-42d7-8f05-92a4f0825f65)

## Algorithmics
Sometimes, the reservations may need a reallocation to be able to insert a new one, which can't be allocated with the current reservations distribution.
For example, in the next reservation distribution, we would not be able to insert a new reservation between days 3 and 10, although this is possible if we reallocate all of them
![Reservation distribution](https://github.com/lucastomic/bookingAPI/assets/65186233/71ff2d40-895c-42f9-9576-a237f8b7f1ed)

To achieve this, we will apply a Backtracking algorithm, where each node is the assignment of a reservation in a different stateroom.
![Backtracking](https://github.com/lucastomic/bookingAPI/assets/65186233/7f9f57fd-a496-435f-8a98-7fa527648858)


The implementation of the algorithm is in the file `internals/reservesReallocator/reallocator.go` with the next methods:
```
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
```
```
func recursiveRealloaction(
	success *bool,
	stateRooms *[]domain.StateRoom,
	reservations *datastructure.Queue[*domain.Reservation],
) {
	if reservations.IsEmpty() {
		*success = true
	} else {
		exploreChildNodes(success, stateRooms, reservations)
	}
}
```
```
func exploreChildNodes(
	success *bool,
	stateRooms *[]domain.StateRoom,
	reservations *datastructure.Queue[*domain.Reservation],
) {
	i := 0
	reservation, _ := reservations.Pop()
	for !*success && len(*stateRooms) > i {
		oldStateRoomId := reservation.StateRoomId()
		reservation.SetStateRoomId((*stateRooms)[i].Id())
		if err := (*stateRooms)[i].AddReservation(*reservation); err == nil {
			recursiveRealloaction(success, stateRooms, reservations)
		}
		if !*success {
			reservation.SetStateRoomId(oldStateRoomId)
			(*stateRooms)[i].RemoveReservation(*reservation)
		}
		i++
	}
	if !*success {
		reservations.Push(reservation)
	}
}
```

## Data structures
For the implementation of the last algorithm, we make use of a Queue, implemented as
