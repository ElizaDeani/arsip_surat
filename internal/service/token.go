package service

import (
	"context"
	"fmt"

	"github.com/ElizaDeani/archivio/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateAccessToken(ctx context.Context, claims entity.JWTCustomClaims) (string, error)
}

type tokenService struct {
	secretKey string
}

func NewTokenService(secretKey string) TokenService {
	return &tokenService{secretKey}
}

func (s *tokenService) GenerateAccessToken(ctx context.Context, claims entity.JWTCustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	fmt.Println("🟢 Token yang dihasilkan:", encodedToken)
	return encodedToken, nil
}
