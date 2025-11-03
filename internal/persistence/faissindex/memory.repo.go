package faissindex

import (
	"context"
	"memory/internal/embedding"
	"memory/internal/persistence"
)

/*
IDEA - folder per conversationID, for each op read the folder, index, write it back
If the conversation is not present, then create its files

what does
*/

type MemoryRepo struct {
	faissClient     *persistence.Client
	embeddingClient embedding.Client
}

func NewMemoryRepo(faissClient *persistence.Client) persistence.VectorMemoryRepo {
	return &MemoryRepo{
		faissClient: faissClient,
	}
}

func (r *MemoryRepo) Index(ctx context.Context, conversationID string, memoryID int, query string) error {
	embedding, err := r.embeddingClient.EmbedOne(ctx, query)
	if err != nil {
		return err
	}
	err = r.faissClient.Index(ctx, conversationID, memoryID, embedding)
	if err != nil {
		return err
	}
	return nil
}

func (r *MemoryRepo) Search(ctx context.Context, conversationID string, query string, topK int) ([]persistence.VectorMemory, error) {
	embedding, err := r.embeddingClient.EmbedOne(ctx, query)
	if err != nil {
		return nil, err
	}
	res, err := r.faissClient.Search(ctx, conversationID, embedding, topK)

	if err != nil {
		return nil, err
	}

	var memories []persistence.VectorMemory
	for id := range res.Labels {
		memories = append(memories, persistence.VectorMemory{
			ID: id,
		})
	}

	return memories, nil
}
