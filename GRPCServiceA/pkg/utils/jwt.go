package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
)

var (
	secretKey = "secret"
	issuer    = "localhost:8000/"
)

func NewToken(name string) string {
	signingKey := []byte(secretKey)
	claims := ent.Claims{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(5) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}
