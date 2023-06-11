package services

import (
	"errors"

	"github.com/golang-jwt/jwt"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/enviroment"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	databaseport.UserRepository
}

func NewAuthenticationService(repo databaseport.UserRepository) authenticationService {
	return authenticationService{repo}
}

func (as authenticationService) Register(email string, rawPassword string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user := domain.NewUserWithoutBoats(email, string(hashPassword))
	err = as.Save(user)
	return err
}

func (as authenticationService) Login(email string, password string) (string, error) {
	err := as.validateEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}
	signedToken, err := as.getSignedToken(email)
	if err != nil {
		return "", errors.New("error singing JWT")
	}
	return signedToken, nil
}

func (as authenticationService) validateEmailAndPassword(email string, password string) error {
	user, err := as.FindById(email)
	if err != nil {
		return exceptions.WrongEmailLogin
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		return exceptions.WrongPasswordLogin
	}
	return nil

}

func (as authenticationService) getSignedToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"sub": email,
			"exp": enviroment.GetJWTExpirationTime(),
		})
	key := enviroment.GetSigningKey()
	signedToken, err := token.SignedString(key)
	return signedToken, err
}

func (as authenticationService) Validate() error {
	return nil
}
