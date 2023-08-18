package reservationtest

import (
	"fmt"
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
)

var mergeTests = []struct {
	r1       domain.Reservation
	r2       domain.Reservation
	expected domain.Reservation
}{
	{
		testutils.GetReservation(true, 2, 3),
		testutils.GetReservation(true, 3),
		testutils.GetReservation(true, 8),
	},
	{
		testutils.GetReservation(true, 2, 3),
		testutils.GetReservation(true, 4),
		testutils.GetReservation(true, 9),
	},
	{
		testutils.GetReservation(true, 3),
		testutils.GetReservation(true, 8),
		testutils.GetReservation(true, 3),
	},
	{
		testutils.GetReservation(false, 3),
		testutils.GetReservation(true, 3),
		testutils.GetReservation(false, 3),
	},
}

func TestMerge(t *testing.T) {
	for i, tt := range mergeTests {
		t.Run(fmt.Sprintf("Test N%v", i), func(t *testing.T) {
			tt.r1.Merge(tt.r2)
			if !tt.r1.Equals(tt.expected) {
				t.Errorf("it's not equals with expected")
			}
		})
	}
}
