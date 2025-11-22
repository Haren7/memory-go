package memory

import (
	"context"
	"fmt"
	"memory/internal/persistence"
	"memory/internal/summarizer"
	"time"

	"github.com/google/uuid"
)

type SemanticService struct {
	vectorMemoryRepo  persistence.VectorMemoryRepoInterface
	rdbmsMemoryRepo   persistence.MemoryRepoInterface
	converstionRepo   persistence.ConversationRepoInterface
	summarizerService summarizer.ServiceInterface
}

func NewSemanticService(
	vectorMemoryRepo persistence.VectorMemoryRepoInterface,
	rdbmsMemoryRepo persistence.MemoryRepoInterface,
	converstionRepo persistence.ConversationRepoInterface,
	summarizerService summarizer.ServiceInterface,
) SemanticServiceInterface {
	return &SemanticService{
		vectorMemoryRepo:  vectorMemoryRepo,
		rdbmsMemoryRepo:   rdbmsMemoryRepo,
		converstionRepo:   converstionRepo,
		summarizerService: summarizerService,
	}
}

func (r *SemanticService) Store(ctx context.Context, conversationID uuid.UUID, query, response string) (uuid.UUID, error) {
	_, err := r.converstionRepo.FetchOne(ctx, conversationID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("semantic: error conversation does not exist, %w", err)
	}
	memoryUUID, err := uuid.NewUUID()
	createdAt := time.Now()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("semantic: error creating memory id, %w", err)
	}
	summarizedResponse, err := r.summarizerService.Summarize(ctx, response)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("semantic: error summarizing response, %w", err)
	}

	_, err = r.rdbmsMemoryRepo.InsertOne(ctx, conversationID, memoryUUID, query, summarizedResponse, createdAt)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("semantic: error persisting memory, %w", err)
	}
	_, err = r.vectorMemoryRepo.Index(ctx, conversationID, memoryUUID, query, summarizedResponse, createdAt)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("semantic: error indexing memory, %w", err)
	}
	return memoryUUID, nil
}

func (r *SemanticService) Retrieve(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error) {
	_, err := r.converstionRepo.FetchOne(ctx, conversationID)
	if err != nil {
		return nil, fmt.Errorf("semantic: error conversation does not exist, %w", err)
	}
	rdbmsMemories, err := r.rdbmsMemoryRepo.FetchManyByConversationID(ctx, conversationID, lastK)
	if err != nil {
		return nil, fmt.Errorf("semantic: error fetching memories, %w", err)
	}
	var memories []Memory
	for _, memory := range rdbmsMemories {
		memories = append(memories, Memory{
			ID:        memory.UUID,
			Query:     memory.Query,
			Response:  memory.Response,
			CreatedAt: memory.CreatedAt,
		})
	}

	return memories, nil
}

func (r *SemanticService) RetrieveSimilar(ctx context.Context, conversationID uuid.UUID, query string, topK int) ([]Memory, error) {
	_, err := r.converstionRepo.FetchOne(ctx, conversationID)
	if err != nil {
		return nil, fmt.Errorf("semantic: error conversation does not exist, %w", err)
	}
	vectorMemories, err := r.vectorMemoryRepo.Search(ctx, conversationID, query, topK)
	if err != nil {
		return nil, fmt.Errorf("semantic: error searching for memories, %w", err)
	}
	var memories []Memory
	for _, memory := range vectorMemories {
		memories = append(memories, Memory{
			ID:        memory.ID,
			Query:     memory.Query,
			Response:  memory.Response,
			CreatedAt: memory.CreatedAt,
		})
	}
	return memories, nil
}
