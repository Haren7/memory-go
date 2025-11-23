package clients

import (
	"context"

	"github.com/haren7/minimal-memory/types"
)

type ShortTermMemoryClient interface {
	Store(ctx context.Context, input types.StoreShortTermMemoryInput) (types.StoreShortTermMemoryOutput, error)
	Retrieve(ctx context.Context, input types.RetrieveShortTermMemoryInput) (types.RetrieveShortTermMemoryOutput, error)
	RegisterConversation(ctx context.Context, input types.RegisterConversationInput) (types.RegisterConversationOutput, error)
}

type SemanticMemoryClient interface {
	Store(ctx context.Context, input types.StoreSemanticMemoryInput) (types.StoreSemanticMemoryOutput, error)
	Retrieve(ctx context.Context, input types.RetrieveSemanticMemoryInput) (types.RetrieveSemanticMemoryOutput, error)
	RegisterConversation(ctx context.Context, input types.RegisterConversationInput) (types.RegisterConversationOutput, error)
}
