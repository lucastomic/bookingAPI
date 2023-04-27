package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type IBoatRepository interface {
	repository[domain.Boat, int]
}
