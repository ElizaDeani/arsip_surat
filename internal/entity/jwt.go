package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

func NewJWTCustomClaims(username, role string, id int64) JWTCustomClaims {
	return JWTCustomClaims{
		Username: username,
		Role:     role,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expire dalam 24 jam
			Issuer:    "archivio",
		},
	}
}
