package embedding

import "context"

type ServiceInterface interface {
	EmbedOne(ctx context.Context, text string) (Embedding, error)
	EmbedMany(ctx context.Context, texts []string) ([]Embedding, error)
}
