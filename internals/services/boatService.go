package services

import (
	"errors"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	reservesreallocator "github.com/lucastomic/naturalYSalvajeRent/internals/reservesReallocator"
	authenticationstate "github.com/lucastomic/naturalYSalvajeRent/internals/state/authentication"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

type boatService struct {
	databaseport.IBoatRepository
	reservationRepo databaseport.IReservationRepository
}

func NewBoatService(
	repo databaseport.IBoatRepository,
	reservationRepo databaseport.IReservationRepository,
) *boatService {
	return &boatService{repo, reservationRepo}
}

func (b boatService) CreateBoat(boat domain.Boat) (domain.Boat, error) {
	if boat.Name() == "" {
		return *domain.EmptyBoat(), errors.New("boat must have a name")
	}
	_, err := b.UpdateBoat(&boat)
	if err != nil {
		return *domain.EmptyBoat(), err
	}
	return boat, nil
}

func (b boatService) UpdateBoat(boat *domain.Boat) (domain.Boat, error) {
	err := b.Save(boat)
	if err != nil {
		return *domain.EmptyBoat(), err
	}
	return *boat, nil
}

func (b boatService) DeleteBoat(boat domain.Boat) error {
	return b.Remove(boat)
}

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

func (b boatService) GetAllBoats() ([]*domain.Boat, error) {
	boats := authenticationstate.UserAuthenticated().Boats()
	return boats, nil
}

func (b boatService) ReservateStaterooms(
	boat domain.Boat,
	reservation domain.Reservation,
	stateroomsNeeded int,
) error {
	if boat.HasDisponibilityFor(reservation, stateroomsNeeded) {
		err := boat.ReservateStaterooms(&reservation, 1)
		if err != nil {
			return exceptions.NewApiError(400, err.Error())
		}
	} else {
		return exceptions.ReservationCollides
	}
	err := b.Save(&boat)
	return err
}

func (b boatService) ReservateStateroomWithReallocation(
	boat domain.Boat,
	reservation domain.Reservation,
) error {
	if boat.HasDisponibilityFor(reservation, 1) {
		err := boat.ReservateStaterooms(&reservation, 1)
		if err != nil {
			return exceptions.ReservationCollides
		}
	} else {
		err := reservesreallocator.RealloacteReserves(&boat, &reservation)
		if err != nil {
			return exceptions.ReservationCollides
		}
	}
	err := b.Save(&boat)
	return err
}

func (b boatService) GetNotEmptyDays(boat domain.Boat) []string {
	notEmptyDays := boat.GetNotEmptyDays()
	return b.parseTimeSliceToString(notEmptyDays)
}

func (b boatService) GetNotAvailableDaysForSharedReservation(
	boat domain.Boat,
	passengers int,
) []string {
	days := boat.GetNotAvailableDaysForSharedReservation(passengers)
	return b.parseTimeSliceToString(days)
}

func (b boatService) GetNotAvailableDaysForCloseReservation(
	boat domain.Boat,
	stateroomsNeeded int,
) []string {
	days := boat.GetNotAvailableDaysForCloseReservation(stateroomsNeeded)
	return b.parseTimeSliceToString(days)
}

func (b boatService) ResevateFullBoat(boat domain.Boat, reservation domain.Reservation) error {
	if boat.HasDisponibilityForEntireBoat(reservation) {
		boat.ReservateEveryStateroom(&reservation)
	} else {
		return exceptions.ReservationCollides
	}
	err := b.Save(&boat)
	return err
}

func (b boatService) parseTimeSliceToString(s []timesimplified.Time) []string {
	var response []string
	for _, day := range s {
		response = append(response, day.ToString())
	}
	return response
}
