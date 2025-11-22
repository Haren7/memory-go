package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MemoryRepoInterface interface {
	SetOne(ctx context.Context, conversationID uuid.UUID, memoryID uuid.UUID, query string, response string, createdAt time.Time) error
	DeleteLastN(ctx context.Context, conversationID uuid.UUID, lastN int) error
	Get(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error)
	Len(ctx context.Context, convesationID uuid.UUID) (int, error)
}

type InMemMemoryRepo struct {
	memories map[uuid.UUID][]Memory
}

func NewInMemMemoryRepo() MemoryRepoInterface {
	return &InMemMemoryRepo{
		memories: make(map[uuid.UUID][]Memory),
	}
}

func (r *InMemMemoryRepo) SetOne(ctx context.Context, conversationID uuid.UUID, memoryID uuid.UUID, query string, response string, createdAt time.Time) error {
	values, exists := r.memories[conversationID]
	if exists {
		values = append(values, Memory{
			ID:        memoryID,
			Query:     query,
			Response:  response,
			CreatedAt: createdAt,
		})
	} else {
		r.memories[conversationID] = []Memory{Memory{ID: memoryID, Query: query, Response: response}}
	}
	return nil
}

func (r *InMemMemoryRepo) DeleteLastN(ctx context.Context, conversationID uuid.UUID, lastN int) error {
	return nil
}

func (r *InMemMemoryRepo) Get(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error) {
	memories, exists := r.memories[conversationID]
	if !exists {
		return nil, fmt.Errorf("error no memories")
	}
	return memories, nil
}

func (r *InMemMemoryRepo) Len(ctx context.Context, convesationID uuid.UUID) (int, error) {
	memories, exists := r.memories[convesationID]
	if !exists {
		return 0, fmt.Errorf("error no memories")
	}
	return len(memories), nil
}
