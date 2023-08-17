package reservationtest

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

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
}
