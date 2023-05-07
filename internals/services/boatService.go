package services

import (
	"errors"
	"time"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	reservesreallocator "github.com/lucastomic/naturalYSalvajeRent/internals/reservesReallocator"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timeParser"
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
		return *domain.EmtyBoat(), errors.New("boat must have a name")
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

// GetFullCapacityDays get a slice of days when all the boat's staterooms are reserved
func (b boatService) GetFullCapacityDays(boat domain.Boat) []string {
	var response []string
	var daysHash map[string]int = make(map[string]int)
	for _, stateRoom := range boat.StateRooms() {
		for _, reservation := range stateRoom.Reservations() {
			reservation.ForEachDay(func(date time.Time) {
				b.updateHashDays(&daysHash, &response, date, boat)
			})

		}
	}
	return response
}

// updateHashDays takes a date and inserts it in the given hash map. If it already exists, it increments its position,
// if it doesn't is inserted with a value of 1. If any date get the same value as the amount of staterooms in the given boat,
// it inserts this date as a string in a string slice specified as parameter
func (b boatService) updateHashDays(daysHash *map[string]int, response *[]string, date time.Time, boat domain.Boat) {
	if _, ok := (*daysHash)[timeParser.ToString(date)]; ok {
		(*daysHash)[timeParser.ToString(date)]++
		if (*daysHash)[timeParser.ToString(date)] == len(boat.StateRooms()) {
			*response = append(*response, timeParser.ToString(date))
		}
	} else {
		(*daysHash)[timeParser.ToString(date)] = 1
	}
}

// AddReservation adds a new reservation to a boat.
// It looks for a free date's range in all the boat's stateRooms which matchs with the reservation one
// If there isn't a free range it reallocates all the reservations (except those which have already started)
// in a way the new reservation can be placed.
// If is impossilbe to allocate the reservation it throws an error.
func (b boatService) AddReservation(boat domain.Boat, reservation domain.Reservation) error {
	couldReserve := false
	for _, stateRoom := range boat.StateRooms() {
		if err := stateRoom.AddReservation(reservation); err == nil {
			couldReserve = true
		}
	}
	if !couldReserve {
		err := reservesreallocator.RealloacteReserves(&boat, &reservation)
		if err != nil {
			return err
		}
	}
	return nil
}
