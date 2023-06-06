package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type BoatRepository interface {
	repository[domain.Boat, int]
	FindByUser(email string) ([]domain.Boat, error)
}
