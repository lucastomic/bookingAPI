package serviceports

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type AuthenticationService interface {
	Register(string, string) error
	// Login returns the correspondient JWT if success
	Login(string, string) (string, error)
	GetUser(email string) (domain.User, error)
}
