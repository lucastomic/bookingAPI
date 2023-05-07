package serviceports

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type IBoatService interface {
	CreateBoat(domain.Boat) (domain.Boat, error)
	UpdateBoat(boat domain.Boat) (domain.Boat, error)
	DeleteBoat(boat domain.Boat) error
	GetBoat(boatId int) (domain.Boat, error)
	GetAllBoats() ([]domain.Boat, error)
	GetFullCapacityDays(domain.Boat) []string
}
