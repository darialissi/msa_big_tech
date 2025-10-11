package utils

import (
	"time"
	"context"

	"github.com/golang-jwt/jwt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
)

func GenerateJWT(ctx context.Context, userId dto.UserID) (*models.Auth, error) {
	secretKey, ok := ctx.Value("jwtSecret").(string)
	if !ok {
		return nil, ErrSecretNotFound
	}

	signingKey := []byte(secretKey)

	// Access token
	accessClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    "auth_service",
		Subject:   string(userId),
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
		Subject:   string(userId),
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

