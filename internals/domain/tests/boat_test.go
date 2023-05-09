package domaintests

import (
	"testing"
	"time"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var today = time.Now()

func newReservation(startDay time.Duration, finalDay time.Duration) *domain.Reservation {
	return domain.NewReservation(0, user1, today.Add(time.Hour*24*startDay), today.Add(time.Hour*24*finalDay), 0, 2)
}

var reserveFromTodayUntil3 = newReservation(0, 3)
var reserveFrom2Until4 = newReservation(2, 4)
var reserveFromTodayUntil4 = newReservation(0, 4)
var reserveFrom1Until4 = newReservation(1, 4)
var reserveFrom8Until9 = newReservation(8, 9)

var boat1 = domain.NewBoat("Test 1", []domain.StateRoom{
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*reserveFromTodayUntil3,
		*reserveFrom8Until9,
	}),
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*reserveFrom2Until4,
		*reserveFrom8Until9,
	}),
})

var boat2 = domain.NewBoat("Test 2", []domain.StateRoom{
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*reserveFromTodayUntil3,
	}),
})

var boat3 = domain.NewBoat("Test 3", []domain.StateRoom{
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*reserveFromTodayUntil4,
	}),
	*domain.NewStateRoom(0, 0, []domain.Reservation{
		*reserveFromTodayUntil3,
		*reserveFrom1Until4,
	}),
})
var boat4 = domain.NewBoat("Test 4", []domain.StateRoom{
	*domain.NewStateRoom(0, 0, []domain.Reservation{}),
})

var getUnstartedReservationsTests = []struct {
	boat     domain.Boat
	expected []*domain.Reservation
}{
	{
		*boat1,
		[]*domain.Reservation{reserveFrom8Until9, reserveFrom2Until4, reserveFrom8Until9},
	},
	{
		*boat2,
		[]*domain.Reservation{},
	},
	{
		*boat3,
		[]*domain.Reservation{reserveFrom1Until4},
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
