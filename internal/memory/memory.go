package memory

import (
	"context"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	Store(ctx context.Context, conversationID uuid.UUID, query, response string) (uuid.UUID, error)
	Retrieve(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error)
}

type SemanticServiceInterface interface {
	Store(ctx context.Context, convesationID uuid.UUID, query, response string) (uuid.UUID, error)
	Retrieve(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error)
	RetrieveSimilar(ctx context.Context, conversationID uuid.UUID, query string, topK int) ([]Memory, error)
}
