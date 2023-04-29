package services

import (
	"errors"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
)

// boatService is a service that provides operations related to boats.
type boatService struct {
	databaseport.IBoatRepository
}

// Returns a new boat service given its repository
func NewBoatService(repo databaseport.IBoatRepository) *boatService {
	return &boatService{repo}
}

// CreateBoat creates a new boat by calling the UpdateBoat() method with the given boat,
// and returns the updated boat or an error if the update fails.
func (b boatService) CreateBoat(boat domain.Boat) (domain.Boat, error) {
	if boat.Name() == "" {
		return *domain.EmtyBoat(), errors.New("Boat must have a name")
	}
	return b.UpdateBoat(boat)
}

// UpdateBoat updates an existing boat by calling the Save() method with the given boat,
// and returns the updated boat or an error if the save operation fails.
func (b boatService) UpdateBoat(boat domain.Boat) (domain.Boat, error) {
	err := b.Save(boat)
	if err != nil {
		return *domain.EmtyBoat(), err
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
		return *domain.EmtyBoat(), err
	}
	if boat.Name() == "" {
		return *domain.EmtyBoat(), exceptions.NotFound
	}
	return boat, nil
}

// GetAllBoats retrieves all boats by calling the FindAll() method,
// and returns a slice of domain.Boat and an error.
func (b boatService) GetAllBoats() ([]domain.Boat, error) {
	boats, err := b.FindAll()
	if err != nil {
		return []domain.Boat{}, err
	}
	return boats, nil
}
