package reservationtest

import (
	"fmt"
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var mergeTests = []struct {
	r1       domain.Reservation
	r2       domain.Reservation
	expected domain.Reservation
}{
	{
		getReservation(true, 2, 3),
		getReservation(true, 3),
		getReservation(true, 8),
	},
	{
		getReservation(true, 2, 3),
		getReservation(true, 4),
		getReservation(true, 9),
	},
	{
		getReservation(true, 3),
		getReservation(true, 8),
		getReservation(true, 3),
	},
	{
		getReservation(false, 3),
		getReservation(true, 3),
		getReservation(false, 3),
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
