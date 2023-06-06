package databaseport

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type UserRepository interface {
	repository[domain.User, string]
}
