package jwtservice

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (claim jwtClaim) IsExpired() bool {
	return claim.ExpiresAt < time.Now().Local().Unix()
}
