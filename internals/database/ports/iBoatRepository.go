package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IBoatRepository interface {
	Save(domain.Boat) error
	FindById(...int) (domain.Boat, error)
}
