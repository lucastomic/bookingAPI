package domaintests

import (
	"strconv"
	"testing"
	"time"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

var user1 = domain.NewUser("Lucas Tomic", "1234212")
var date = time.Date(2023, 05, 03, 20, 34, 58, 651387237, time.UTC)

var reservation2Days = domain.NewReservation(0, user1, date, date.Add(time.Hour*24*2), 0, 0)

var containsTests = []struct {
	reservation1 domain.Reservation
	date         time.Time
	expected     bool
}{
	{
		*reservation2Days,
		date.Add(time.Hour * 24),
		true,
	},
	{
		*reservation2Days,
		date.Add(time.Hour * 72),
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
