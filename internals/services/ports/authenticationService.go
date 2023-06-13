package serviceports

type AuthenticationService interface {
	Register(string, string) error
	// Login returns the correspondient JWT if success
	Login(string, string) (string, error)
	Validate(string) error
}
