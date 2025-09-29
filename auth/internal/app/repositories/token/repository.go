package token

import (
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db map[string]string // key-value
}

func NewRepository() *Repository {
	return &Repository{
		db: make(map[string]string),
	}
}