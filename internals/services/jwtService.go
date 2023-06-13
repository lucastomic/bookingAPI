package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lucastomic/naturalYSalvajeRent/internals/enviroment"
)

type jwtService struct {
}

func NewJWTService() jwtService {
	return jwtService{}
}

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (service jwtService) GenerateToken(email string) (string, error) {
	//TODO: This Signing method must be changed for a asymetric one, such as SigningMethodES256
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + enviroment.GetJWTExpirationTime(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return service.signToken(token)
}

func (service jwtService) signToken(token *jwt.Token) (string, error) {
	key := enviroment.GetSigningKey()
	return token.SignedString([]byte(key))
}

func (as authenticationService) Validate(signedToken string) (err error) {
	jwtKey := enviroment.GetSigningKey()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return
}
