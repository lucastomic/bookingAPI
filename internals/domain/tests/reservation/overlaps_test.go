package reservationtest

import (
	"fmt"
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
)

var overlapsTests = []struct {
	r1       domain.Reservation
	r2       domain.Reservation
	expected bool
}{
	{
		testutils.GetReservationInDays(0, 1),
		testutils.GetReservationInDays(2, 3),
		false,
	},
	{
		testutils.GetReservationInDays(0, 1),
		testutils.GetReservationInDays(1, 3),
		true,
	},
	{
		testutils.GetReservationInDays(0, 0),
		testutils.GetReservationInDays(0, 1),
		true,
	},
	{
		testutils.GetReservationInDays(3, 4),
		testutils.GetReservationInDays(1, 5),
		true,
	},
	{
		testutils.GetReservationInDays(1, 5),
		testutils.GetReservationInDays(3, 4),
		true,
	},
	{
		testutils.GetReservationInDays(0, 1),
		testutils.GetReservationInDays(0, 1),
		true,
	},
	{
		testutils.GetReservationInDays(0, 0),
		testutils.GetReservationInDays(0, 0),
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
