package stateroomtests

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
)

var canReservateTests = []struct {
	name        string
	stateRoom   domain.StateRoom
	reservation *domain.Reservation
	expected    bool
}{
	{
		"should be true, because should add reservation in day 5",
		testutils.GetStateRoom([]*domain.Reservation{
			testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(5, 5, true, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(7, 8, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(10, 12, false, 2),
		}),
		testutils.GetReservationWithDaysOpenAndPassengers(5, 5, true, 3),
		true,
	},
	{
		"should be false, because reservation in day 5 is close and can't merge",
		testutils.GetStateRoom([]*domain.Reservation{
			testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(5, 5, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(7, 8, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(10, 12, false, 2),
		}),
		testutils.GetReservationWithDaysOpenAndPassengers(5, 5, true, 3),
		false,
	},
	{
		"should be false, because reservations dates doesn't match",
		testutils.GetStateRoom([]*domain.Reservation{
			testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(5, 6, true, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(7, 8, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(10, 12, false, 2),
		}),
		testutils.GetReservationWithDaysOpenAndPassengers(5, 5, true, 3),
		false,
	},
	{
		"should be true, because dates are free",
		testutils.GetStateRoom([]*domain.Reservation{
			testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(7, 8, false, 3),
			testutils.GetReservationWithDaysOpenAndPassengers(10, 12, false, 2),
		}),
		testutils.GetReservationWithDaysOpenAndPassengers(5, 6, false, 3),
		true,
	},
	{
		"should be true, because there isn't any reservation",
		testutils.GetStateRoom([]*domain.Reservation{}),
		testutils.GetReservationWithDaysOpenAndPassengers(5, 6, false, 3),
		true,
	},
}

func TestCanReservate(t *testing.T) {
	for _, tt := range canReservateTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stateRoom.CanReservate(*tt.reservation)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
