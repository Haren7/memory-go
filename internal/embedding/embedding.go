package embedding

import "context"

type Embedding struct {
	Dim    int
	Vector []float32
}

type Client interface {
	EmbedOne(ctx context.Context, text string) (Embedding, error)
	EmbedMany(ctx context.Context, texts []string) ([]Embedding, error)
}
