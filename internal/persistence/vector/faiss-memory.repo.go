package vector

import (
	"context"
	"memory/internal/embedding"
	"memory/internal/persistence"
	"time"

	"github.com/google/uuid"
)

type FaissMemoryRepo struct {
	faissClient     *FaissClient
	embeddingClient embedding.ServiceInterface
	rdbmsMemoryRepo persistence.MemoryRepoInterface
}

func NewFaissMemoryRepo(faissClient *FaissClient) persistence.VectorMemoryRepoInterface {
	return &FaissMemoryRepo{
		faissClient: faissClient,
	}
}

func (r *FaissMemoryRepo) Index(ctx context.Context, conversationID, memoryID uuid.UUID, query, response string, createdAt time.Time) (persistence.VectorMemory, error) {
	memoryId, err := r.rdbmsMemoryRepo.InsertOne(ctx, conversationID, memoryID, query, response, createdAt)
	if err != nil {
		return persistence.VectorMemory{}, err
	}
	embedding, err := r.embeddingClient.EmbedOne(ctx, query)
	if err != nil {
		return persistence.VectorMemory{}, err
	}
	err = r.faissClient.Index(ctx, conversationID.String(), memoryId, embedding)
	if err != nil {
		return persistence.VectorMemory{}, err
	}
	return persistence.VectorMemory{
		ID:       uuid.New(),
		Query:    query,
		Response: response,
	}, nil
}

func (r *FaissMemoryRepo) Search(ctx context.Context, conversationID uuid.UUID, query string, topK int) ([]persistence.VectorMemory, error) {
	embedding, err := r.embeddingClient.EmbedOne(ctx, query)
	if err != nil {
		return nil, err
	}
	faissResponse, err := r.faissClient.Search(ctx, conversationID.String(), embedding, topK)
	var memoryIds []int
	for _, id := range faissResponse.Labels {
		memoryIds = append(memoryIds, int(id))
	}
	rdbmsMemories, err := r.rdbmsMemoryRepo.FetchMany(ctx, memoryIds)
	if err != nil {
		return nil, err
	}
	var vectorMemories []persistence.VectorMemory
	for _, memory := range rdbmsMemories {
		vectorMemories = append(vectorMemories, persistence.VectorMemory{
			ID:        memory.UUID,
			Query:     memory.Query,
			Response:  memory.Response,
			CreatedAt: memory.CreatedAt,
		})
	}
	return vectorMemories, nil
}
