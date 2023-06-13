package serviceports

type JWTService interface {
	GenerateToken(email string) (string, error)
	Validate(string) error
}
