package services

import (
	"errors"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	reservesreallocator "github.com/lucastomic/naturalYSalvajeRent/internals/reservesReallocator"
	authenticationstate "github.com/lucastomic/naturalYSalvajeRent/internals/state/authentication"
)

// boatService is a service that provides operations related to boats.
type boatService struct {
	databaseport.BoatRepository
	reservationRepo databaseport.IReservationRepository
}

// Returns a new boat service given its repository
func NewBoatService(repo databaseport.BoatRepository, reservationRepo databaseport.IReservationRepository) *boatService {
	return &boatService{repo, reservationRepo}
}

// CreateBoat creates a new boat by calling the UpdateBoat() method with the given boat,
// and returns the updated boat or an error if the update fails.
// TODO: it returns a wrong ID
func (b boatService) CreateBoat(boat domain.Boat) (domain.Boat, error) {
	if boat.Name() == "" {
		return *domain.EmptyBoat(), errors.New("boat must have a name")
	}
	_, err := b.UpdateBoat(boat)
	if err != nil {
		return *domain.EmptyBoat(), err
	}
	return boat, nil
}

// UpdateBoat updates an existing boat by calling the Save() method with the given boat,
// and returns the updated boat or an error if the save operation fails.
func (b boatService) UpdateBoat(boat domain.Boat) (domain.Boat, error) {
	err := b.Save(boat)
	if err != nil {
		return *domain.EmptyBoat(), err
	}
	return boat, nil
}

// DeleteBoat deletes a boat by calling the Remove() method with the given boat,
// and returns an error if the removal operation fails.
func (b boatService) DeleteBoat(boat domain.Boat) error {
	err := b.Remove(boat)
	if err != nil {
		return err
	}
	return nil
}

// GetBoat retrieves a boat by its ID by calling the FindById() method with the given boat ID,
// and returns the found boat or an error if the boat is not found.
func (b boatService) GetBoat(boatId int) (domain.Boat, error) {
	boat, err := b.FindById(boatId)
	if err != nil {
		return *domain.EmptyBoat(), err
	}
	if boat.Name() == "" {
		return *domain.EmptyBoat(), exceptions.NotFound
	}
	return boat, nil
}

// GetAllBoats retrieves all boats by calling the FindAll() method,
// and returns a slice of domain.Boat and an error.
func (b boatService) GetAllBoats() ([]domain.Boat, error) {
	boats := authenticationstate.UserAuthenticated().Boats()
	// if err != nil {
	// 	return []domain.Boat{}, err
	// }
	return boats, nil
}

// GetFullCapacityDays get a slice of days when all the boat's staterooms are reserved
func (b boatService) GetFullCapacityDays(boat domain.Boat) []string {
	var response []string
	var days = boat.GetFullCapacityDays()
	for _, day := range days {
		response = append(response, day.ToString())
	}
	return response
}

// AddReservation adds a new reservation to a boat.
// It looks for a free date's range in all the boat's stateRooms which matchs with the reservation one
// If there isn't a free range it reallocates all the reservations (except those which have already started)
// in a way the new reservation can be placed.
// If is impossilbe to allocate the reservation it throws an error.
func (b boatService) AddReservation(boat domain.Boat, reservation domain.Reservation) error {
	couldReserve := boat.AddReservation(&reservation)
	if !couldReserve {
		err := reservesreallocator.RealloacteReserves(&boat, &reservation)
		if err != nil {
			return exceptions.ReservationCollides
		}
	}
	err := b.Save(boat)
	return err
}

// GetNotEmptyDays retrives those days where there is at least one reservation of a boat.
func (b boatService) GetNotEmptyDays(boat domain.Boat) []string {
	var response []string
	var days = boat.GetNotEmptyDays()
	for _, day := range days {
		response = append(response, day.ToString())
	}
	return response
}

// ReservateFullBoat reservates all the staterooms in the boat.
// Returns true if the reservation was allocated propperly and false if there is no free range for the reservation
func (b boatService) ResevateFullBoat(boat domain.Boat, reservation domain.Reservation) error {
	couldReservate := boat.ReservateFullBoat(&reservation)
	if !couldReservate {
		return exceptions.ReservationCollides
	}
	err := b.Save(boat)
	return err
}
