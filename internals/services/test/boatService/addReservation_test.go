package boatservicetests

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

// Dates for testing
var today = timesimplified.Now()

func newReservation(startDay int, finalDay int) *domain.Reservation {
	return domain.NewReservation(0, user1, today.AddDays(startDay), today.AddDays(finalDay), 0, 2)
}

var stateRoom1 = domain.NewStateRoom(0, 0, []domain.Reservation{
	*newReservation(0, 3), *newReservation(5, 8), *newReservation(10, 11), *newReservation(13, 16),
})

var stateRoom2 = domain.NewStateRoom(0, 0, []domain.Reservation{
	*newReservation(0, 1), *newReservation(3, 7), *newReservation(9, 11), *newReservation(13, 14),
})

var stateRoom3 = domain.NewStateRoom(0, 0, []domain.Reservation{
	*newReservation(1, 4), *newReservation(7, 10), *newReservation(14, 17),
})

var stateRoom4 = domain.NewStateRoom(0, 0, []domain.Reservation{
	*newReservation(0, 2), *newReservation(4, 7), *newReservation(11, 13),
})

var boatWithReservations = domain.NewBoat("Boat with reservations", []domain.StateRoom{
	*stateRoom1,
	*stateRoom2,
	*stateRoom3,
	*stateRoom4,
})

var emptyBoat = domain.NewBoat("Empty boat", []domain.StateRoom{
	*domain.EmptyStateRoom(),
})

var firstDaysReserved = domain.NewBoat("First days reserved", []domain.StateRoom{
	*stateRoom1,
})

var possibleBoat = domain.NewBoat("Unable reservate", []domain.StateRoom{
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*newReservation(0, 3),
		*newReservation(4, 7),
		*newReservation(11, 17),
	}),
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*newReservation(3, 5),
		*newReservation(8, 9),
		*newReservation(0, 1),
	}),
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*newReservation(1, 3),
		*newReservation(6, 9),
	}),
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*newReservation(0, 5),
		*newReservation(7, 10),
	}),
})

var addReservationTests = []struct {
	domain.Boat
	domain.Reservation
	expected bool //true if is expected to be able to reallocate the reservations
}{
	{
		*boatWithReservations,
		*newReservation(8, 17),
		true,
	},
	{
		*emptyBoat,
		*newReservation(1, 4),
		true,
	},
	{
		*firstDaysReserved,
		*newReservation(0, 4),
		false,
	},
	{
		*possibleBoat,
		*newReservation(4, 10),
		true,
	},
}

// TODO: REPLACE BOATSERVICE WITH ONE WITH MOCKED REPO
// THE TEST WON'T BE RIGHT UNTIL IS DONE
func TestAddReservation(t *testing.T) {
	for _, tt := range addReservationTests {
		t.Run(tt.Name(), func(t *testing.T) {
			err := boatService.AddReservation(tt.Boat, tt.Reservation)
			if (err != nil) == tt.expected {
				t.Errorf("Failed")
			}

		})
	}
}
