package persistence

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ConversationRepoInterface interface {
	FetchOne(ctx context.Context, conversationID uuid.UUID) (Conversation, error)
	InsertOne(ctx context.Context, agent, user string, conversationID uuid.UUID, createdAt time.Time) (int, error)
}

type MemoryRepoInterface interface {
	FetchOne(ctx context.Context, conversationID uuid.UUID) (Memory, error)
	FetchMany(ctx context.Context, memoryIds []int) ([]Memory, error)
	FetchManyByConversationID(ctx context.Context, conversationID uuid.UUID, limit int) ([]Memory, error)
	InsertOne(ctx context.Context, conversationID uuid.UUID, memoryID uuid.UUID, query, response string, createdAt time.Time) (int, error)
}

type VectorMemoryRepoInterface interface {
	Index(ctx context.Context, conversationID uuid.UUID, memoryID uuid.UUID, query, response string, createdAt time.Time) (VectorMemory, error)
	Search(ctx context.Context, conversationID uuid.UUID, query string, topK int) ([]VectorMemory, error)
}
