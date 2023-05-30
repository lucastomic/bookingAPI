package boatservicetests

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

var date = timesimplified.NewTime(2023, 05, 3)

func newRes(from int, to int) domain.Reservation {
	return *domain.NewReservation(0, user1, date.AddDays(from), date.AddDays(to), 0, 1)
}

// StateRooms for testing
var stateRoom2daysReserved = *domain.NewStateRoom(0, 0, []domain.Reservation{
	newRes(0, 1),
})

var stateRoomReserved3days = *domain.NewStateRoom(1, 0, []domain.Reservation{
	newRes(0, 2),
})

var stateRoomReserved4days = *domain.NewStateRoom(2, 0, []domain.Reservation{
	newRes(0, 3),
})

var stateRoomReserved2daysAnd4th = *domain.NewStateRoom(0, 0, []domain.Reservation{
	newRes(0, 1),
	newRes(3, 3),
})

var boatWith3days = domain.NewBoatWithId(0, "3 days", []domain.StateRoom{
	stateRoom2daysReserved,
	stateRoomReserved3days,
	stateRoomReserved4days,
})

var boatWithoutDays = domain.NewBoatWithId(0, "Without days", []domain.StateRoom{
	stateRoom2daysReserved,
	*domain.NewStateRoom(1, 0, []domain.Reservation{
		*domain.NewReservation(0, user1, date.AddDays(3), date.AddDays(4), 0, 1),
	}),
})

var boatWithSeparatedRanges = domain.NewBoatWithId(2, "Separated range days", []domain.StateRoom{
	stateRoomReserved4days,
	stateRoomReserved2daysAnd4th,
	stateRoomReserved2daysAnd4th,
	stateRoomReserved2daysAnd4th,
})
var fullCapacityDaysTest = []struct {
	boat     domain.Boat
	expected []string
}{
	{*boatWith3days, []string{"2023-05-03", "2023-05-04"}},
	{*boatWithoutDays, []string{}},
	{*boatWithSeparatedRanges, []string{"2023-05-03", "2023-05-04", "2023-05-06"}},
}

func TestGetFullCapacityDays(t *testing.T) {
	for _, tt := range fullCapacityDaysTest {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := boatService.GetFullCapacityDays(tt.boat)
			if !compareSlices(got, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
func compareSlices(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
