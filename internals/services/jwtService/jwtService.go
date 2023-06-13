package jwtservice

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

func (as jwtService) Validate(signedToken string) (err error) {
	token, err := as.parseStringToToken(signedToken)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*jwtClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.IsExpired() {
		return errors.New("token expired")
	}
	return nil
}

func (service jwtService) GenerateToken(email string) (string, error) {
	claims := service.getClaims(email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return service.signToken(token)
}

func (service jwtService) parseStringToToken(token string) (*jwt.Token, error) {
	jwtKey := enviroment.GetSigningKey()
	return jwt.ParseWithClaims(
		token,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
}
func (service jwtService) getClaims(email string) *jwtClaim {
	//TODO: This Signing method must be changed for a asymetric one, such as SigningMethodES256
	return &jwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: service.getJWTExpirationTime(),
		},
	}
}

func (service jwtService) signToken(token *jwt.Token) (string, error) {
	key := enviroment.GetSigningKey()
	return token.SignedString([]byte(key))
}

func (service jwtService) getJWTExpirationTime() int64 {
	return time.Now().Unix() + enviroment.GetJWTExpirationTime()
}
