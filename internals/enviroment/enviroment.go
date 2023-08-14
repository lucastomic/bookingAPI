package enviroment

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	signingKey                      = "SINGING_KEY"
	jwtExpirationTimeInMilliseconds = "EXPIRATION_TIME_MILLISECONDS"
	databaseUser                    = "MYSQL_USER"
	databasePassword                = "MYSQL_PASSWORD"
	databaseName                    = "MYSQL_DATABASE"
	sgbd                            = "SGBD"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		print("Error initalizing .env variables")
	}
}

func GetSigningKey() string {
	return os.Getenv(signingKey)
}
func GetJWTExpirationTime() int64 {
	var rawResponse string = os.Getenv(jwtExpirationTimeInMilliseconds)
	responseParsed, _ := strconv.Atoi(rawResponse)
	return int64(responseParsed)
}
func GetDatabaseUser() string {
	return os.Getenv(databaseUser)
}
func GetDatabasePassword() string {
	return os.Getenv(databasePassword)
}
func GetDatabaseName() string {
	return os.Getenv(databaseName)
}
func GetSGBD() string {
	return os.Getenv(sgbd)
}
