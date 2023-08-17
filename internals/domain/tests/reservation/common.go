package reservationtest

import (
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
func getReservationInDays(starterDay int, finalDay int) domain.Reservation {
	return *domain.NewReservationWithoutClient(
		0,
		getDayFromToday(starterDay),
		getDayFromToday(finalDay),
		false,
		0,
	)
}

func getReservationWithDaysOpenAndPassengers(starterDay int, finalDay int, isOpen bool, passengers ...int) domain.Reservation {
	res := domain.NewReservationWithoutClient(
		0,
		getDayFromToday(starterDay),
		getDayFromToday(finalDay),
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
