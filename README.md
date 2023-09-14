# Booking API
Booking API is a simple API for reservation management in a boat with different staterooms. Supports the next features:

-  Booking the full boat for N days
-  Booking just N staterooms of the boat
-  Booking the boat in shared way. This mean, other client's can join to your reservation.


# API usage
This section won't get in details deeply (explanation of the parameters, possible returns, etc) because it's not the porpouse of the repository, but the explanation of the Software engeneering areas involved in its construction.

The application counts with the next endpoints:
### `POST /auth/register`
Creates a new user, expects a JSON body request like this:
```json
{
  "email": "myEmail@gmail.com",
  "password":"secretPassword"
}
```
### `POST /auth/login`
Authenticate the user and retrieve the Json Web Token, expects a JSON body request like this:
```json
{
  "email": "myEmail@gmail.com",
  "password":"secretPassword"
}
```
### `POST /boat`
Creates a new boat, expects a JSON body request like this:
```json
{
  "name": "BoatName",
  "maxCapacity":10
}
```
| Nombre del Parámetro | Tipo   | Descripción                            |
| --------------------- | ------ | --------------------------------------|
| `name`                | String | Boat's name                  |
| `maxCapacity`         | Number | Boat's max capacity      |

If successful it will return a JSON like the following

```json
{
    "id": 14,
    "maxCapacity": 10,
    "name": "BoatName",
    "owner": "myEmail@gmail.com",
    "stateRoom": null
}
```
### `PUT /stateRoom/add/:boatId`
Adds a new stateroom to a boat, given the boat ID

### `GET /boat/:id`
Get a specific boat, given its ID. As the Create-Boat endpoint, returns a JSON like the following
```json
{
    "id": 14,
    "maxCapacity": 10,
    "name": "BoatName",
    "owner": "myEmail@gmail.com",
    "stateRoom": null
}
```

### `GET /boat/reserved/:id`
Get the days when a boat it's in its full capacity (this means those days when all the staterooms are reserved) given its ID

### `GET /boat/notEmpty/:id`
Get those days where there is at least one reservation of a boat given its ID

### `POST /boat/reservate`
Makes a reservation in a boat, expects a JSON body request like this:
```json
{
	"email":"ltomicb@gmail.com",
	"phone":"623029321",
	"firstDay":"2023-12-06",
	"lastDay":"2023-12-09",
	"boatId":13
}
```
This endpoint could return an error if there is not enough space for the new reservation
### `POST /boat/reservateFullBoat`
Reservates reserves the entire boat for the specified reservation. Expects a JSON body request like this:
```json
{
	"email":"ltomicb@gmail.com",
	"phone":"623029321",
	"firstDay":"2023-12-06",
	"lastDay":"2023-12-09",
	"boatId":13
}
```
This endpoint could return an error if there is not enough space for the new reservation
### `DELETE /boat/:id`
Deletes a specific boat, given its ID

### `PUT /stateRoom/add/:boatId`
Adds a new stateroom to a boat, given the boat ID

### `DELETE /reservation/:id`
Deletes a specific reservation, given its ID

## Algorithmics
Sometimes, the reservations may need a reallocation to be able to insert a new one, which can't be allocated with the current reservations distribution.
For example, in the next reservation distribution, we would not be able to insert a new reservation between days 3 and 10, although this is possible if we reallocate all of them
![Reservation distribution](https://github.com/lucastomic/bookingAPI/assets/65186233/71ff2d40-895c-42f9-9576-a237f8b7f1ed)

To achieve this, we will apply a Backtracking algorithm, where each node is the assignment of a reservation in a different stateroom.
![Backtracking](https://github.com/lucastomic/bookingAPI/assets/65186233/7f9f57fd-a496-435f-8a98-7fa527648858)


The implementation of the algorithm is in the file `internals/reservesReallocator/reallocator.go` with the next methods:
(comments ommited)

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
The algorithm takes into account that the reservations which have already started, can't be reallocated (the clients could be already in their staterooms)

