package domaintests

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

var today = timesimplified.Now()

func newReservation(startDay int, finalDay int) *domain.Reservation {
	return domain.NewReservation(0, today.AddDays(startDay), today.AddDays(finalDay), user1, false, 0, 0)
}

var boat1 = domain.NewBoat("Test 1", []*domain.StateRoom{
	domain.NewStateRoom(0, 0, []*domain.Reservation{
		newReservation(0, 3),
		newReservation(8, 9),
	}),
	domain.NewStateRoom(0, 0, []*domain.Reservation{
		newReservation(2, 4),
		newReservation(8, 9),
	}),
}, "")

var boat2 = domain.NewBoat("Test 2", []*domain.StateRoom{
	domain.NewStateRoom(0, 0, []*domain.Reservation{
		newReservation(0, 3),
	}),
}, "")

var boat3 = domain.NewBoat("Test 3", []*domain.StateRoom{
	domain.NewStateRoom(0, 0, []*domain.Reservation{
		newReservation(0, 4),
	}),
	domain.NewStateRoom(0, 0, []*domain.Reservation{
		newReservation(0, 3),
		newReservation(1, 4),
	}),
}, "")

var boat4 = domain.NewBoat("Test 4", []*domain.StateRoom{
	domain.NewStateRoom(0, 0, []*domain.Reservation{}),
}, "")

var getUnstartedReservationsTests = []struct {
	boat     domain.Boat
	expected []*domain.Reservation
}{
	{
		*boat1,
		[]*domain.Reservation{newReservation(8, 9), newReservation(2, 4), newReservation(8, 9)},
	},
	{
		*boat2,
		[]*domain.Reservation{},
	},
	{
		*boat3,
		[]*domain.Reservation{newReservation(1, 4)},
	},
	{
		*boat4,
		[]*domain.Reservation{},
	},
}

func TestGetUnestartedReservations(t *testing.T) {
	for _, tt := range getUnstartedReservationsTests {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := tt.boat.GetUnstartedReservations()
			if !compareSlices(got, tt.expected) {
				t.Errorf("Expected: \n%v, got: \n%v", sliceToString(tt.expected), sliceToString(got))
			}

		})
	}
}

func sliceToString(s []*domain.Reservation) string {
	var response string
	for _, res := range s {
		response += res.String()
	}
	return response
}

func compareSlices(s1 []*domain.Reservation, s2 []*domain.Reservation) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i].FirstDay() != s2[i].FirstDay() || s1[i].LastDay() != s2[i].LastDay() {
			return false
		}
	}
	return true
}
