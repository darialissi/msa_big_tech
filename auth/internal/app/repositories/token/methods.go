package token

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)

func (r *Repository) Save(ctx context.Context, auth *dto.AuthRefresh) (error) {
	id := string(auth.ID)

	r.mu.Lock()
	defer r.mu.Unlock()

	r.kv[id] = auth.RefreshToken

	return nil
}

func (r *Repository) FetchById(ctx context.Context, userId models.UserID) (string, error) {
	id := string(userId)

	r.mu.RLock()
	defer r.mu.RUnlock()

	value, exists := r.kv[id]

	if !exists {
		return "", fmt.Errorf("No refresh token found userId=%s", userId)
	}

	return value, nil
}