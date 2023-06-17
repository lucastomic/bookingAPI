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
	claims, err := as.getClaims(token)
	if err != nil {
		return err
	}
	if claims.IsExpired() {
		return errors.New("token expired")
	}
	return nil
}

func (service jwtService) GenerateToken(email string) (string, error) {
	claims := service.generateClaims(email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return service.signToken(token)
}

func (service jwtService) GetEmail(token string) (string, error) {
	tokenParsed, err := service.parseStringToToken(token)
	if err != nil {
		return "", errors.New("invalid token string")
	}
	claims, err := service.getClaims(tokenParsed)
	if err != nil {
		return "", err
	}
	return claims.Email, nil
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
func (service jwtService) generateClaims(email string) *jwtClaim {
	//TODO: This Signing method must be changed for a asymetric one, such as SigningMethodES256
	return &jwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: service.getJWTExpirationTime(),
		},
	}
}

func (service jwtService) getClaims(token *jwt.Token) (jwtClaim, error) {
	claims, ok := token.Claims.(*jwtClaim)
	if !ok {
		return jwtClaim{}, errors.New("couldn't parse claims")
	}
	return *claims, nil
}

func (service jwtService) signToken(token *jwt.Token) (string, error) {
	key := enviroment.GetSigningKey()
	return token.SignedString([]byte(key))
}

func (service jwtService) getJWTExpirationTime() int64 {
	return time.Now().Unix() + enviroment.GetJWTExpirationTime()
}
