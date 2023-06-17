package authenticationstate

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

var userAuthenticated domain.User

func SetAuthenticatedUser(user domain.User) {
	userAuthenticated = user
}
func RemoveAuthenticatedUser() {
	userAuthenticated = domain.User{}
}

func UserAuthenticated() domain.User {
	return userAuthenticated
}

func IsAuthenticated() bool {
	return userAuthenticated.Email() != ""
}
