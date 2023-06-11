package enviroment

import "os"

const (
	signingKey                      = "SINGING_KEY"
	jwtExpirationTimeInMilliseconds = "EXPIRATION_TIME_MILLISECONDS"
)

func GetSigningKey() string {
	return os.Getenv(signingKey)
}
func GetJWTExpirationTime() string {
	return os.Getenv(jwtExpirationTimeInMilliseconds)
}
