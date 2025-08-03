// Package jwt ....
package jwt

import (
	"time"

	"api_gin_bun/config"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Email  string `json:"email"`
	Nombre string `json:"nombre"`
	ID     int64  `json:"id"`
	jwt.RegisteredClaims
}

func JWTGenerator(email, nombre string, id int64) (string, error) {
	claims := MyCustomClaims{
		Email:  email,
		Nombre: nombre,
		ID:     id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.CfgEnv.SecretJWT))
	return tokenString, err
}
