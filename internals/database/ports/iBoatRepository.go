package databaseport

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type IBoatRepository interface {
	Repository[domain.Boat, int]
	FindByUser(email string) ([]*domain.Boat, error)
}
