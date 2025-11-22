package vector

import (
	"context"
	"fmt"
	"memory/internal/embedding"
	"memory/internal/persistence"
	"time"

	"github.com/google/uuid"
	"github.com/philippgille/chromem-go"
)

type memory struct {
	UUID           uuid.UUID
	Query          string
	Respones       string
	ConversationID uuid.UUID
	CreatedAt      time.Time
}

type ChromemMemoryRepo struct {
	db               *chromem.DB
	embeddingService embedding.ServiceInterface
}

func NewChromemMemoryRepo(db *chromem.DB, embeddingService embedding.ServiceInterface) persistence.VectorMemoryRepoInterface {
	return &ChromemMemoryRepo{
		db:               db,
		embeddingService: embeddingService,
	}
}

func (r *ChromemMemoryRepo) Index(ctx context.Context, conversationID, memoryID uuid.UUID, query, response string, createdAt time.Time) (persistence.VectorMemory, error) {
	collection, err := r.db.GetOrCreateCollection(conversationID.String(), nil, nil)
	if err != nil {
		return persistence.VectorMemory{}, fmt.Errorf("chromem: error getting or creating collection, %w", err)
	}
	embedding, err := r.embeddingService.EmbedOne(ctx, query)
	if err != nil {
		return persistence.VectorMemory{}, fmt.Errorf("chromem: error embedding query, %w", err)
	}
	document := chromem.Document{
		ID:        memoryID.String(),
		Embedding: embedding.Vector,
		Metadata: r.transformToMap(memory{
			UUID:           memoryID,
			Query:          query,
			Respones:       response,
			ConversationID: conversationID,
			CreatedAt:      createdAt,
		}),
	}
	err = collection.AddDocument(ctx, document)
	if err != nil {
		return persistence.VectorMemory{}, fmt.Errorf("chromem: error adding document, %w", err)
	}
	return persistence.VectorMemory{
		ID:             memoryID,
		ConversationID: conversationID,
		Query:          query,
		Response:       response,
		CreatedAt:      createdAt,
	}, nil
}

func (r *ChromemMemoryRepo) Search(ctx context.Context, conversationID uuid.UUID, query string, topK int) ([]persistence.VectorMemory, error) {
	embedding, err := r.embeddingService.EmbedOne(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("chromem: error embedding query, %w", err)
	}
	collection := r.db.GetCollection(conversationID.String(), nil)
	count := collection.Count()
	if count < topK {
		topK = count
	}
	result, err := collection.QueryEmbedding(ctx, embedding.Vector, topK, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("chromem: error querying embedding, %w", err)
	}
	var memories []memory
	for _, result := range result {
		memory, err := r.transformFromMap(result.Metadata)
		if err != nil {
			return nil, fmt.Errorf("chromem: error transforming from map, %w", err)
		}
		memories = append(memories, memory)
	}
	var vectorMemories []persistence.VectorMemory
	for _, memory := range memories {
		vectorMemories = append(vectorMemories, persistence.VectorMemory{
			ID:             memory.UUID,
			ConversationID: memory.ConversationID,
			Query:          memory.Query,
			Response:       memory.Respones,
			CreatedAt:      memory.CreatedAt,
		})
	}
	return vectorMemories, nil
}

func (r *ChromemMemoryRepo) transformToMap(data memory) map[string]string {
	return map[string]string{
		"uuid":           data.UUID.String(),
		"query":          data.Query,
		"response":       data.Respones,
		"conversationId": data.ConversationID.String(),
		"createdAt":      data.CreatedAt.Format(time.RFC3339),
	}
}

func (r *ChromemMemoryRepo) transformFromMap(data map[string]string) (memory, error) {
	createdAt, err := time.Parse(time.RFC3339, data["createdAt"])
	if err != nil {
		return memory{}, fmt.Errorf("chromem: error parsing created at, %w", err)
	}
	return memory{
		Query:          data["query"],
		Respones:       data["response"],
		ConversationID: uuid.MustParse(data["conversationId"]),
		CreatedAt:      createdAt,
	}, nil
}
