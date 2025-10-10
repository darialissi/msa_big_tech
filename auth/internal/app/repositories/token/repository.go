package token

import (
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	kv map[string]string // key-value
}

func NewRepository() *Repository {
	return &Repository{
		kv: make(map[string]string),
	}
}