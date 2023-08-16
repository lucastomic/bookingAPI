package serviceports

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// IBoatService is the port for using the BoatService
type IBoatService interface {
	// CreateBoat creates a new boat.
	// It takes a domain.Boat object as input and returns the created domain.Boat object along with any error encountered.
	CreateBoat(domain.Boat) (domain.Boat, error)
	// UpdateBoat updates an existing boat.
	// It takes a domain.Boat object as input and returns the updated domain.Boat object along with any error encountered.
	UpdateBoat(boat *domain.Boat) (domain.Boat, error)
	// DeleteBoat deletes a boat.
	// It takes a domain.Boat object as input and returns an error if any occurred.
	DeleteBoat(boat domain.Boat) error
	// GetBoat retrieves a boat by its ID.
	// It takes a boatId (integer) as input and returns the corresponding domain.Boat object along with any error encountered.
	GetBoat(boatId int) (domain.Boat, error)
	// GetAllBoats retrieves all the boats.
	// It returns a slice of domain.Boat objects along with any error encountered.
	GetAllBoats() ([]*domain.Boat, error)
	// GetFullCapacityDays returns a list of dates on which the specified boat is fully booked.
	// It takes a domain.Boat object as input and returns a slice of strings representing the full capacity days.
	GetFullCapacityDays(domain.Boat) []string
	// GetNotEmptyDays returns a list of dates on which the specified boat has at least one reservation.
	// It takes a domain.Boat object as input and returns a slice of strings representing the non-empty days.
	GetNotEmptyDays(domain.Boat) []string
	// AddReservation adds a new reservation to the specified boat.
	// It takes a domain.Boat object and a domain.Reservation object as input and returns an error if any occurred.
	// For example, it would return an error if the reservation can't be allocated because there is not enough space.
	ReservateStateroom(domain.Boat, domain.Reservation) error
	// ResevateFullBoat reserves the entire boat for the specified reservation.
	// It takes a domain.Boat object and a domain.Reservation object as input and returns an error if any occurred.
	// For example, it would return an error if the reservation can't be allocated because there is not enough space.
	ResevateFullBoat(boat domain.Boat, reservation domain.Reservation) error
}
