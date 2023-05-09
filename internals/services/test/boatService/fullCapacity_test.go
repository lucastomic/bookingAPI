package boatservicetests

import (
	"testing"
	"time"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var date = time.Date(2023, 05, 03, 20, 34, 58, 651387237, time.UTC)

var reservation2days = *domain.NewReservation(0, user1, date, date.Add(time.Hour*48), 0, 1)
var reservation3Days = *domain.NewReservation(0, user1, date, date.Add(time.Hour*72), 0, 1)
var reservation4Days = *domain.NewReservation(0, user1, date, date.Add(time.Hour*96), 0, 2)

// StateRooms for testing
var stateRoom2daysReserved = *domain.NewStateRoom(0, 0, []domain.Reservation{
	reservation2days,
})

var stateRoomReserved3days = *domain.NewStateRoom(1, 0, []domain.Reservation{
	reservation3Days,
})

var stateRoomReserved4days = *domain.NewStateRoom(2, 0, []domain.Reservation{
	reservation4Days,
})

var stateRoomReserved2daysAnd4th = *domain.NewStateRoom(0, 0, []domain.Reservation{
	reservation2days,
	*domain.NewReservation(0, user1, date.Add(time.Hour*72), date.Add(time.Hour*96), 0, 1),
})

var boatWith2days = domain.NewBoatWithId(0, "2 days", []domain.StateRoom{
	stateRoom2daysReserved,
	stateRoomReserved3days,
	stateRoomReserved4days,
})

var boatWithoutDays = domain.NewBoatWithId(0, "Without days", []domain.StateRoom{
	stateRoom2daysReserved,
	*domain.NewStateRoom(1, 0, []domain.Reservation{
		*domain.NewReservation(0, user1, date.Add(time.Hour*48), date.Add(time.Hour*72), 0, 1),
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
	{*boatWith2days, []string{"2023-05-03", "2023-05-04"}},
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
