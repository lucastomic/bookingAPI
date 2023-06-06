package domaintests

import (
	"strconv"
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

var user1 = domain.NewClient("Lucas Tomic", "1234212")
var date = timesimplified.NewTime(2023, 05, 06)

var reservation2Days = domain.NewReservation(0, user1, date, date.AddDays(2), 0, 0)

var containsTests = []struct {
	reservation1 domain.Reservation
	date         timesimplified.Time
	expected     bool
}{
	{
		*reservation2Days,
		date.AddDays(1),
		true,
	},
	{
		*reservation2Days,
		date.AddDays(3),
		false,
	},
	{
		*reservation2Days,
		date,
		true,
	},
}

func TestContains(t *testing.T) {
	for i, tt := range containsTests {
		t.Run("Test N: "+strconv.Itoa(i), func(t *testing.T) {
			got := tt.reservation1.Contains(tt.date)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}

		})
	}
}
