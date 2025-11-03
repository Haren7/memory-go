package persistence

import (
	"context"
)

type RdbmsConversationRepo interface {
	FetchOne(ctx context.Context, agent string, user string) (RdbmsConversation, error)
	InsertOne(ctx context.Context, agent string, user string) (RdbmsConversation, error)
}

type RdbmsMemoryRepo interface {
	FetchOne(ctx context.Context, conversationID string) (RdbmsMemory, error)
	InsertOne(ctx context.Context, conversationID string, query string, response string) (RdbmsMemory, error)
}

type VectorMemoryRepo interface {
	Index(ctx context.Context, conversationID string, memoryID int, query string) error
	Search(ctx context.Context, conversationID string, query string, topK int) ([]VectorMemory, error)
}
