package databaseport

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type IUserRepository interface {
	Repository[domain.User, string]
}
