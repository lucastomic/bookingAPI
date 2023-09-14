package boattest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

var getNotAvailableDaysForStateroomsTests = []struct {
	boat             domain.Boat
	stateroomsNeeded int
	expected         []timesimplified.Time
}{
	{
		testutils.GetBoat(
			"Not enoguh reservations for 0-3 because two st are already reserved",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
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
		},
	},
	{
		testutils.GetBoat(
			"Close reservations mustn't collide with shared ones",
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
		1,
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
			"Enough space always",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, false, 2),
			},
			[]*domain.Reservation{},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
			},
		),
		2,
		[]timesimplified.Time{},
	},
	{
		testutils.GetBoat(
			"Must return full days",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, false),
			},
		),
		1,
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
			"Must return not empty days",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, false),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
			},
		),
		3,
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
			"Close and shared mustn't collide",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, true, 2),
			},
			[]*domain.Reservation{},
			[]*domain.Reservation{},
		),
		4,
		[]timesimplified.Time{
			testutils.GetDay(0),
			testutils.GetDay(1),
			testutils.GetDay(2),
			testutils.GetDay(3),
		},
	},
}

func TestGetNotAvailableDaysForStaterooms(t *testing.T) {
	for _, tt := range getNotAvailableDaysForStateroomsTests {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := tt.boat.GetNotAvailableDaysForCloseReservation(tt.stateroomsNeeded)
			if !compareTimeSlices(got, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
