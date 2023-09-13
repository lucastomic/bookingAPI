package boattest

import (
	"testing"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	testutils "github.com/lucastomic/naturalYSalvajeRent/internals/testing/utils"
)

var hasDisponibilityTestSet = []struct {
	boat            domain.Boat
	stateroomNeeded int
	reservation     domain.Reservation
	expected        bool
}{
	{
		testutils.GetBoat(
			"must be true becuase has enough free saterooms",
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
		true,
	},
	{
		testutils.GetBoat(
			"must be false becuase are just 2 free staterooms in this dates",
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
		false,
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
		false,
	},
	{
		testutils.GetBoat(
			"must be true becuase new reservation can be merged",
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
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
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
		true,
	},
	{
		testutils.GetBoat(
			"must be false becuase the current reservation is close",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
		),
		1,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
		false,
	},
	{
		testutils.GetBoat(
			"must be true becuase can be merged with the shared one",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
			},
		),
		1,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 1, true, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must be true becuase one stateroom is free in this dates",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(4, 9, false, 2),
			},
		),
		1,
		*testutils.GetReservationWithDaysOpenAndPassengers(0, 3, false, 2),
		true,
	},
	{
		testutils.GetBoat(
			"must be false becuase there aren't enough staterooms",
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
			[]*domain.Reservation{
				testutils.GetReservationWithDaysOpenAndPassengers(0, 1, false, 2),
			},
		),
		3,
		*testutils.GetReservationWithDaysOpenAndPassengers(2, 3, false, 2),
		false,
	},
}

func TestReservateFullBoat(t *testing.T) {
	for _, tt := range hasDisponibilityTestSet {
		t.Run(tt.boat.Name(), func(t *testing.T) {
			got := tt.boat.HasDisponibilityFor(tt.reservation, tt.stateroomNeeded)
			if got != tt.expected {
				t.Errorf("Expected: \n%v, got: \n%v", tt.expected, got)
			}
		})
	}
}
