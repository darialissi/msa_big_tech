package token

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)

func (r *Repository) Save(auth *dto.AuthRefresh) (error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.db[auth.ID.String()] = auth.RefreshToken

	return nil
}

func (r *Repository) Fetch(userId dto.UserID) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	value, exists := r.db[userId.String()]

	if !exists {
		return "", fmt.Errorf("No refresh token userId=%s", userId.String())
	}

	return value, nil
}