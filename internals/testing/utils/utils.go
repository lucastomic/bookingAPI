package testutils

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

func GetBoat(name string, reservations ...[]*domain.Reservation) domain.Boat {
	var stateRooms []*domain.StateRoom
	for _, reservation := range reservations {
		stateRooms = append(stateRooms, GetStateRoom(reservation))
	}
	return *domain.NewBoat(
		name,
		stateRooms,
		"",
	)
}

func GetStateRoom(reservations []*domain.Reservation) *domain.StateRoom {
	return domain.NewStateRoom(
		0,
		0,
		reservations,
	)
}

func GetDayFromToday(daysAfter int) timesimplified.Time {
	return timesimplified.Now().AddDays(daysAfter)
}

func GetClientWithPassengers(passengers int) *domain.Client {
	return domain.NewClient("Jhon Doe", "123321123", passengers)
}

func GetReservation(isOpen bool, passengers ...int) domain.Reservation {
	return createReservation(0, 0, isOpen, passengers...)
}

func GetReservationInDays(starterDay, finalDay int) domain.Reservation {
	return createReservation(starterDay, finalDay, false)
}

func GetReservationWithDaysOpenAndPassengers(starterDay, finalDay int, isOpen bool, passengers ...int) *domain.Reservation {
	res := createReservation(starterDay, finalDay, isOpen, passengers...)
	return &res
}

func createReservation(starterDay, finalDay int, isOpen bool, passengers ...int) domain.Reservation {
	res := domain.NewReservationWithoutClient(
		0,
		GetDayFromToday(starterDay),
		GetDayFromToday(finalDay),
		isOpen,
		0,
	)
	res.SetMaxCapacity(10)
	clients := make([]*domain.Client, 0, len(passengers))
	for _, n := range passengers {
		clients = append(clients, GetClientWithPassengers(n))
	}
	res.SetClients(clients)
	return *res
}
