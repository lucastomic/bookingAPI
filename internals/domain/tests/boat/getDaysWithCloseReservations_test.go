package boattest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

var daysWithCloseReservationsTests = []struct {
	boat     domain.Boat
	expected []timesimplified.Time
}{
	{

		testutils.GetBoat(
			"",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, false),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
			},
		),
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
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, false),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
				testutils.GetReservationWithDaysOpenAndPassengers(7, 7, true),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false),
			},
		),
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
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, true),
			},
			[]*domain.Reservation{},
		),
		[]timesimplified.Time{},
	},
	{

		testutils.GetBoat(
			"",
			[]*domain.Reservation{},
			[]*domain.Reservation{},
			[]*domain.Reservation{},
		),
		[]timesimplified.Time{},
	},
}

func TestGetFullCapacityDays(t *testing.T) {
	for _, tt := range daysWithCloseReservationsTests {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := tt.boat.GetDaysWithCloseReservations()
			if !compareTimeSlices(got, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
func compareTimeSlices(s1 []timesimplified.Time, s2 []timesimplified.Time) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !s1[i].Equals(s2[i]) {
			return false
		}
	}
	return true
}
