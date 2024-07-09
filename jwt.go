package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) createJWT(expiresInSeconds int, userID int) (string, error) {
	expirationTime := 0
	if expiresInSeconds > 0 {
		expirationTime = expiresInSeconds
	}
	mySigningKey := []byte(cfg.JWTSecret)
	claims :=
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expirationTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "chirpy",
			Subject:   fmt.Sprintf("%d", userID),
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}
