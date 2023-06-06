package serviceports

type AuthenticationService interface {
	SignIn() error
	Login() error
	Validate() error
}
