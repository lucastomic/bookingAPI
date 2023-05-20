# Booking API
Booking API is a simple API for reservation management in a boat with different staterooms. This is designed mainly for the porpouse of showing mu backend skills.
This is the stack used for the project development:
- Go
- Docker
- MYSQL
- AWS

# API usage
This section won't get in details deeply (explanation of the parameters, possible returns, etc) because it's not the porpouse of the repository, but the explanation of the Software engeneering areas involved in its construction.

The application counts with the next endpoints:
### `GET /boat/:id`
Get a specific boat, given its ID
### `GET /boat/reserved/:id`
Get the days when a boat it's in its full capacity (this means those days when all the staterooms are reserved) given its ID
### `POST /boat`
Creates a new boat, expects a JSON body request like this:
```json
{
  "name": "Boat 1"
}
```
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
### `DELETE /boat/:id`
Deletes a specific boat, given its ID

### `PUT /stateRoom/add/:boatId`
Adds a new stateroom to a boat, given the boat ID

### `DELETE /reservation/:id`
Deletes a specific reservation, given its ID

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

## Data structures
For the implementation of the last algorithm, we make use of a Queue, implemented in `internals/datastructure/queue.go` as 
(comments ommited)

```
type Queue[T any] struct {
	array []T
}

func (q *Queue[T]) Push(el T) {
	q.array = append(q.array, el)
}

func (q *Queue[T]) Pop() (T, error) {
	if len(q.array) == 0 {
		return *new(T), errors.New("there is no elements in queue")
	}
	response := q.array[0]
	q.array = q.array[1:]
	return response, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.array) == 0
}
func (q Queue[T]) Size() int {
	return len(q.array)
}

func NewQueue[T any](arr []T) *Queue[T] {
	return &Queue[T]{arr}
}

```
## Database
The database is managed with MYSQL and raw SQL (no ORM). The architecture of the database code implementation, can be seen in the [architecture section](#architecture), where the `database` struct is the implementation of the singleton pattern.

Although the structure of the database is quite simple, this is the layout of the diagram.

![database diagram](https://github.com/lucastomic/bookingAPI/assets/65186233/bf933efa-52d7-4332-9377-577aeee3739d)

And the SQL used to generate the schema
```
CREATE SCHEMA IF NOT EXISTS naturalYSalvaje;

USE naturalYSalvaje;

CREATE TABLE boat(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE stateRoom(
  id INT NOT NULL,
  boatId INT NOT NULL,
  PRIMARY KEY(boatId, id),
  FOREIGN KEY(boatId) REFERENCES boat(id)
);

CREATE TABLE reservation(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  firstDay DATE NOT NULL,
  lastDay DATE NOT NULL,
  boatId INT NOT NULL,
  stateRoomId INT NOT NULL,
  
  PRIMARY KEY(id),
  FOREIGN KEY(boatId, stateRoomId) REFERENCES stateRoom(boatId,id)
);
```

## Testing
As this is not a real project, there was no need of spending time in a extensive coverage. However, for the demonstration of knowledge, some functions were covered




```
var user1 = domain.NewUser("Lucas Tomic", "1234212")
var date = time.Date(2023, 05, 03, 20, 34, 58, 651387237, time.UTC)

var reservation2Days = domain.NewReservation(0, user1, date, date.Add(time.Hour*24*2), 0, 0)

var containsTests = []struct {
	reservation1 domain.Reservation
	date         time.Time
	expected     bool
}{
	{
		*reservation2Days,
		date.Add(time.Hour * 24),
		true,
	},
	{
		*reservation2Days,
		date.Add(time.Hour * 72),
		false,
	},
	{
		*reservation2Days,
		date,
		true,
	},
}

func TestContains(t *testing.T) {
	for i, tt := range containsTests {
		t.Run("Test N: "+strconv.Itoa(i), func(t *testing.T) {
			got := tt.reservation1.Contains(tt.date)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}

		})
	}
}
```

## Docker
Also, it makes use of Docker for managing the database and dockerize the application, using besides [air](https://github.com/cosmtrek/air) for live reload.

## Deployment
For the deployment of the project, there where used AWS services, specifically AWS EC2 for the API hosting and RDS for the database.
(The public endpoint for using the API is not exposed as the security section of the project hasn't been finished yet)

## Pending to finish
The development of this project hasn't take into account the security of the API. This area will be covered in the future
