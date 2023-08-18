package boattest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var emptyBoat = domain.NewBoat("empty boat", []*domain.StateRoom{
	domain.NewStateRoom(0, 0, []*domain.Reservation{}),
	domain.NewStateRoom(0, 0, []*domain.Reservation{}),
	domain.NewStateRoom(0, 0, []*domain.Reservation{}),
}, "")

var notAvailable = domain.NewBoat("not available boat", []*domain.StateRoom{
	domain.NewStateRoom(0, 0, []*domain.Reservation{
		newReservation(0, 3),
	}),
	domain.NewStateRoom(1, 0, []*domain.Reservation{}),
	domain.NewStateRoom(2, 0, []*domain.Reservation{}),
}, "")

var reservateFullBoat = []struct {
	boat     domain.Boat
	res      domain.Reservation
	expected bool
}{
	{
		*emptyBoat,
		*newReservation(0, 1),
		true,
	},
	{
		*notAvailable,
		*newReservation(0, 1),
		false,
	},
	{
		*notAvailable,
		*newReservation(4, 9),
		true,
	},
}

func TestReservateFullBoat(t *testing.T) {
	for _, tt := range reservateFullBoat {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := tt.boat.ReservateEveryStateroom(&tt.res)
			if got != tt.expected {
				t.Errorf("Expected: \n%v, got: \n%v", tt.expected, got)
			}

		})
	}
}
