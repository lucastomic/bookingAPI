package reservationtest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
)

var canMergeTests = []struct {
	name        string
	client      domain.Reservation
	reservation domain.Reservation
	expected    bool
}{
	{
		"should be able to add client",
		testutils.GetReservation(true, 2),
		testutils.GetReservation(true, 4),
		true,
	},
	{
		"should reject beceuase exceeds max capacity",
		testutils.GetReservation(true, 2),
		testutils.GetReservation(true, 9),
		false,
	},
	{
		"should reject because resevation is close",
		testutils.GetReservation(true, 2),
		testutils.GetReservation(false, 2),
		false,
	},
	{
		"should reject because are 14 passengers and capacity is 10",
		testutils.GetReservation(true, 4),
		testutils.GetReservation(true, 4, 4, 2),
		false,
	},
	{
		"should reject because are 11 passengers and capacity is 10",
		testutils.GetReservation(true, 2),
		testutils.GetReservation(true, 4, 4, 1),
		false,
	},
	{
		"should be able because there are 10 passengers and the capcity is 10",
		testutils.GetReservation(true, 4),
		testutils.GetReservation(true, 2, 4),
		true,
	},
	{
		"should reject because the dates are not the same, even though boths are one day reservation",
		testutils.GetReservation(true, 4),
		*testutils.GetReservationWithDaysOpenAndPassengers(12, 12, true, 2, 4),
		false,
	},
	{
		"should reject because are different days",
		testutils.GetReservation(true, 4),
		*testutils.GetReservationWithDaysOpenAndPassengers(12, 23, true, 2, 4),
		false,
	},
}

func TestCanMerge(t *testing.T) {
	for _, tt := range canMergeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.reservation.CanMerge(tt.client)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
