package services

import databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"

type authenticationService struct {
	databaseport.UserRepository
}

func (as authenticationService) SignIn(email string, password string) error {
	return nil
}
func (as authenticationService) Login() error {
	return nil
}
func (as authenticationService) Validate() error {
	return nil
}
