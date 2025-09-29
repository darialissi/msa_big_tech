package utils

import (
	"time"
	"os"

	"github.com/golang-jwt/jwt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
)


func getEnv(key string, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
		return value
    }

    return defaultValue
}

func GenerateJWT(userId dto.UserID) (*models.Auth, error) {
	signingKey := []byte(getEnv("JWT_SECRET", ""))

	// Access token
	accessClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    "auth_service",
		Subject:   userId.String(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(signingKey)
	if err != nil {
		return nil, ErrAccessToken
	}

	// Refresh token
	refreshClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		Issuer:    "auth_service",
		Subject:   userId.String(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(signingKey)
	if err != nil {
		return nil, ErrRefreshToken
	}

	return &models.Auth{
		ID: models.UserID(userId),
		Token: models.Token{
			AccessToken:  signedAccessToken,
			RefreshToken: signedRefreshToken,
		},
	}, nil
}

