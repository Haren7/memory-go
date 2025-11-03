package embedding

import "context"

type OpenAI struct {
}

func NewOpenAI() Client {
	return &OpenAI{}
}

func (r *OpenAI) EmbedOne(ctx context.Context, text string) (Embedding, error) {
	return Embedding{}, nil
}

func (r *OpenAI) EmbedMany(ctx context.Context, texts []string) ([]Embedding, error) {
	return []Embedding{}, nil
}
