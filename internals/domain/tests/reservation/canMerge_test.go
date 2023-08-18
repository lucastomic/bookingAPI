package reservationtest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var canMergeTests = []struct {
	name        string
	client      domain.Reservation
	reservation domain.Reservation
	expected    bool
}{
	{
		"should be able to add client",
		getReservation(true, 2),
		getReservation(true, 4),
		true,
	},
	{
		"should reject beceuase exceeds max capacity",
		getReservation(true, 2),
		getReservation(true, 9),
		false,
	},
	{
		"should reject because resevation is close",
		getReservation(true, 2),
		getReservation(false, 2),
		false,
	},
	{
		"should reject because are 14 passengers and capacity is 10",
		getReservation(true, 4),
		getReservation(true, 4, 4, 2),
		false,
	},
	{
		"should reject because are 11 passengers and capacity is 10",
		getReservation(true, 2),
		getReservation(true, 4, 4, 1),
		false,
	},
	{
		"should be able because there are 10 passengers and the capcity is 10",
		getReservation(true, 4),
		getReservation(true, 2, 4),
		true,
	},
	{
		"should reject because the dates are not the same, even though boths are one day reservation",
		getReservation(true, 4),
		getReservationWithDaysOpenAndPassengers(12, 12, true, 2, 4),
		false,
	},
	{
		"should reject because are different days",
		getReservation(true, 4),
		getReservationWithDaysOpenAndPassengers(12, 23, true, 2, 4),
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
