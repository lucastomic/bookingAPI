package domain

import (
	"errors"
)

type ReservationManager struct {
	staterooms  []*StateRoom
	maxCapacity int
}

func (b *ReservationManager) HasDisponibilityFor(
	reservation Reservation,
	stateroomsNeeded int,
) bool {
	availableStaterooms := 0
	for _, stateroom := range b.staterooms {
		if stateroom.CanReservate(reservation) {
			availableStaterooms++
		}
	}
	return availableStaterooms >= stateroomsNeeded
}

func (b *ReservationManager) HasDisponibilityForEntireBoat(reservation Reservation) bool {
	return b.HasDisponibilityFor(reservation, len(b.staterooms))
}

func (b *ReservationManager) ReservateStaterooms(reservation *Reservation, staterooms int) error {
	if !b.HasDisponibilityFor(*reservation, staterooms) {
		return errors.New("there is not enough space for this reservation")
	}
	if reservation.isOpen {
		return errors.New(
			"only close reservations can reserve only one stateroom. Shared resrevations must reservate all of them",
		)
	}
	stateroomReserved := 0
	for _, stateRoom := range b.staterooms {
		if (*stateRoom).CanReservate(*reservation) {
			reservation.SetMaxCapacity(b.maxCapacity)
			stateRoom.Reservate(reservation)
			stateroomReserved++
		}
		if stateroomReserved == staterooms {
			break
		}
	}
	return nil
}

func (b *ReservationManager) ReservateStateroom(reservation *Reservation) error {
	return b.ReservateStaterooms(reservation, 1)
}

func (b *ReservationManager) ReservateEveryStateroom(reservation *Reservation) error {
	if !b.HasDisponibilityForEntireBoat(*reservation) {
		return errors.New("there is not enough space for this reservation")
	}
	for _, stateroom := range b.staterooms {
		err := stateroom.Reservate(reservation)
		if err != nil {
			return err
		}
		reservation.SetMaxCapacity(b.maxCapacity)
	}
	return nil
}
