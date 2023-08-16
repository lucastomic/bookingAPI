package services

import (
	"errors"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	databaseport.IUserRepository
	jwtService serviceports.JWTService
}

func NewAuthenticationService(repo databaseport.IUserRepository, jwtService serviceports.JWTService) authenticationService {
	return authenticationService{
		repo,
		jwtService,
	}
}

func (as authenticationService) Register(email string, rawPassword string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user := domain.NewUserWithoutBoats(email, string(hashPassword))
	err = as.Save(&user)
	return err
}

func (as authenticationService) Login(email string, password string) (string, error) {
	err := as.validateEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}
	signedToken, err := as.jwtService.GenerateToken(email)
	if err != nil {
		return "", errors.New("error singing JWT")
	}
	return signedToken, nil
}

func (as authenticationService) GetUser(email string) (domain.User, error) {
	return as.FindById(email)
}

func (as authenticationService) validateEmailAndPassword(email string, rawPassword string) error {
	user, err := as.FindById(email)
	if err != nil {
		return exceptions.WrongEmailLogin
	}
	err = as.compareHashAndPassword(rawPassword, user.Password())
	if err != nil {
		return exceptions.WrongPasswordLogin
	}
	return nil
}

func (as authenticationService) compareHashAndPassword(rawPassword string, hash string) error {
	rawPasswordParsed := []byte(rawPassword)
	hashParsed := []byte(hash)
	return bcrypt.CompareHashAndPassword(hashParsed, rawPasswordParsed)
}
