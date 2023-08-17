package reservationtest

import (
	"fmt"
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var overlapsTests = []struct {
	r1       domain.Reservation
	r2       domain.Reservation
	expected bool
}{
	{
		getReservationInDays(0, 1),
		getReservationInDays(2, 3),
		false,
	},
	{
		getReservationInDays(0, 1),
		getReservationInDays(1, 3),
		true,
	},
	{
		getReservationInDays(0, 0),
		getReservationInDays(0, 1),
		true,
	},
	{
		getReservationInDays(3, 4),
		getReservationInDays(1, 5),
		true,
	},
	{
		getReservationInDays(1, 5),
		getReservationInDays(3, 4),
		true,
	},
	{
		getReservationInDays(0, 1),
		getReservationInDays(0, 1),
		true,
	},
	{
		getReservationInDays(0, 0),
		getReservationInDays(0, 0),
		true,
	},
}

func TestOverlaps(t *testing.T) {
	for i, tt := range overlapsTests {
		t.Run(fmt.Sprintf("Test N%v", i), func(t *testing.T) {
			got := tt.r1.Overlaps(tt.r2)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
