package cache

import (
	"context"
	"memory/internal/persistence"
)

type MemoryCacheInterface interface {
	Set(ctx context.Context, conversationID string, query string, response string) error
	Get(ctx context.Context, conversationID string, lastK int) ([]persistence.CacheMemory, error)
}

type MemoryCache struct {
	size     int
	memories map[string][]persistence.CacheMemory
}

func NewMemoryCache(size int) MemoryCacheInterface {
	return &MemoryCache{
		size:     size,
		memories: make(map[string][]persistence.CacheMemory),
	}
}

func (r *MemoryCache) Set(ctx context.Context, conversationID string, query string, response string) error {
	// if len > size, remove oldest
	return nil
}

func (r *MemoryCache) Get(ctx context.Context, conversationID string, lastK int) ([]persistence.CacheMemory, error) {
	return nil, nil
}
