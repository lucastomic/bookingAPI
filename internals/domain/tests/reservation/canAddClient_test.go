package reservationtest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

func getDayFromToday(daysAfter int) timesimplified.Time {
	return timesimplified.Now().AddDays(daysAfter)
}

func getClientWithPassengers(passengers int) *domain.Client {
	return domain.NewClient("Jhon doe", "123321123", passengers)
}

func getReservation(isOpen bool, passengers ...int) domain.Reservation {
	res := domain.NewReservationWithoutClient(
		0,
		getDayFromToday(0),
		getDayFromToday(0),
		isOpen,
		0,
	)
	res.SetMaxCapacity(10)
	clients := []*domain.Client{}
	for _, n := range passengers {
		clients = append(clients, getClientWithPassengers(n))
	}
	res.SetClients(clients)
	return *res
}

var canAddClientTests = []struct {
	name        string
	client      domain.Client
	reservation domain.Reservation
	expected    bool
}{
	{
		"should be able to add client",
		*getClientWithPassengers(2),
		getReservation(true, 4),
		true,
	},
	{
		"should reject beceuase exceeds max capacity",
		*getClientWithPassengers(2),
		getReservation(true, 9),
		false,
	},
	{
		"should reject because resevation is close",
		*getClientWithPassengers(2),
		getReservation(false, 2),
		false,
	},
	{
		"should reject because are 14 passengers and capacity is 10",
		*getClientWithPassengers(4),
		getReservation(true, 4, 4, 2),
		false,
	},
	{
		"should reject because are 11 passengers and capacity is 10",
		*getClientWithPassengers(2),
		getReservation(true, 4, 4, 1),
		false,
	},
	{
		"should be able because there are 10 passengers and the capcity is 10",
		*getClientWithPassengers(4),
		getReservation(true, 2, 4),
		true,
	},
}

func TestCanAddClient(t *testing.T) {
	for _, tt := range canAddClientTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.reservation.CanAddClient(tt.client)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
