package entities

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Name string
	jwt.StandardClaims
}

func (c Claims) Valid() error {
	return nil
}
