package boattest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
)

var reservasteTestSet = []struct {
	boat            domain.Boat
	stateroomNeeded int
	reservation     domain.Reservation
	expectsError    bool
}{
	{
		testutils.GetBoat(
			"must work because it has enough space and reservation is not shared",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
		),
		3,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
		false,
	},
	{
		testutils.GetBoat(
			"must return error becuase are just 2 free staterooms in this dates",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(1, 2, false, 2),
				testutils.GetReservationWithDaysOpenAndPassengers(3, 6, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
			},
		),
		3,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must be false becuase new reservation is close",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
			},
		),
		4,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must be true becuase new reservation can be merged",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
		),
		1,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must be true becuase all staterooms are free",
			[]*domain.Reservation{},
			[]*domain.Reservation{},
			[]*domain.Reservation{},
			[]*domain.Reservation{},
		),
		2,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
		false,
	},
	{
		testutils.GetBoat(
			"must return error becuase reservation is shared",
			[]*domain.Reservation{},
			[]*domain.Reservation{},
			[]*domain.Reservation{},
			[]*domain.Reservation{},
		),
		2,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must return error becuase there is no enough space",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
		),
		4,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must work becuase one stateroom is free in this dates",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, false, 2),
			},
		),
		1,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
		false,
	},
	{
		testutils.GetBoat(
			"must return error becuase there aren't enough staterooms",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
		),
		3,
		*testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
		true,
	},
}

func TestReservateStateroom(t *testing.T) {
	for _, tt := range reservasteTestSet {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			err := tt.boat.ReservateStaterooms(&tt.reservation, tt.stateroomNeeded)
			if (err != nil) != tt.expectsError {
				t.Errorf("Expected error: \n%v, got: \n%v", tt.expectsError, err)
			}
		})
	}
}

func TestMaxCapacityIsSet(t *testing.T) {
	for _, tt := range reservasteTestSet {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			err := tt.boat.ReservateStaterooms(&tt.reservation, tt.stateroomNeeded)
			if err == nil {
				if tt.reservation.MaxCapacity() != tt.boat.MaxCapacity() {
					t.Errorf(
						"Expected max maxCapacity: \n%v, got: \n%v",
						tt.boat.MaxCapacity(),
						tt.reservation.MaxCapacity(),
					)
				}
			}
		})
	}
}
