package clients

import (
	"context"
	"memory/types"
)

type ShortTermMemoryClientInterface interface {
	Store(ctx context.Context, input types.StoreShortTermMemoryInput) error
	Retrieve(ctx context.Context, input types.RetrieveShortTermMemoryInput) (types.RetrieveShortTermMemoryOutput, error)
	RegisterConversation(ctx context.Context, input types.RegisterConversationInput) (types.RegisterConversationOutput, error)
}

type SemanticMemoryClientInterface interface {
	Store(ctx context.Context, input types.StoreSemanticMemoryInput) error
	Retrieve(ctx context.Context, input types.RetrieveSemanticMemoryInput) (types.RetrieveSemanticMemoryOutput, error)
	RegisterConversation(ctx context.Context, input types.RegisterConversationInput) (types.RegisterConversationOutput, error)
}

type EpisodicMemoryClientInterface interface {
	Store()
	Retrieve()
	RegisterConversation()
}
