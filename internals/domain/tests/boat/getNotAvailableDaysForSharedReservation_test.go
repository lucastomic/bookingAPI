package boattest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

var getNotAvailableDaysForSharedReservationTests = []struct {
	boat       domain.Boat
	passengers int
	expected   []timesimplified.Time
}{
	{
		testutils.GetBoat(
			"",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, false),
			},
			[]*domain.Reservation{},
		),
		2,
		[]timesimplified.Time{
			testutils.GetDay(0),
			testutils.GetDay(1),
			testutils.GetDay(2),
			testutils.GetDay(3),
			testutils.GetDay(7),
		},
	},
	{
		testutils.GetBoat(
			"",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, false),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
			},
		),
		9,
		[]timesimplified.Time{
			testutils.GetDay(0),
			testutils.GetDay(1),
			testutils.GetDay(2),
			testutils.GetDay(3),
			testutils.GetDay(4),
			testutils.GetDay(5),
			testutils.GetDay(6),
			testutils.GetDay(7),
			testutils.GetDay(8),
			testutils.GetDay(9),
		},
	},
	{
		testutils.GetBoat(
			"",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
			},
		),
		2,
		[]timesimplified.Time{},
	},
	{
		testutils.GetBoat(
			"",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, false),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
			},
		),
		2,
		[]timesimplified.Time{
			testutils.GetDay(7),
		},
	},
}

func TestGetNotAvailableDaysForSharedReservations(t *testing.T) {
	for _, tt := range getNotAvailableDaysForSharedReservationTests {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := tt.boat.GetNotAvailableDaysForSharedReservation(tt.passengers)
			if !compareTimeSlices(got, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
