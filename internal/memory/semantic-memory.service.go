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
	summarizerService summarizer.ServiceInterface
}

func NewSemanticService(
	vectorMemoryRepo persistence.VectorMemoryRepoInterface,
	rdbmsMemoryRepo persistence.MemoryRepoInterface,
	summarizerService summarizer.ServiceInterface,
) SemanticServiceInterface {
	return &SemanticService{
		vectorMemoryRepo:  vectorMemoryRepo,
		rdbmsMemoryRepo:   rdbmsMemoryRepo,
		summarizerService: summarizerService,
	}
}

func (r *SemanticService) Store(ctx context.Context, conversationID uuid.UUID, query, response string) (uuid.UUID, error) {
	memoryUUID, err := uuid.NewUUID()
	createdAt := time.Now()
	if err != nil {
		fmt.Printf("\n")
		return uuid.UUID{}, fmt.Errorf("error creating memory id")
	}
	summarizedResponse, err := r.summarizerService.Summarize(ctx, response)
	if err != nil {
		fmt.Printf("\n")
		return uuid.UUID{}, fmt.Errorf("error summarizing response")
	}

	_, err = r.rdbmsMemoryRepo.InsertOne(ctx, conversationID, memoryUUID, query, summarizedResponse, createdAt)
	if err != nil {
		fmt.Printf("\n")
		return uuid.UUID{}, fmt.Errorf("error persisting memory")
	}
	_, err = r.vectorMemoryRepo.Index(ctx, conversationID, memoryUUID, query, summarizedResponse, createdAt)
	if err != nil {
		fmt.Printf("\n")
		return uuid.UUID{}, fmt.Errorf("error indexing memory")
	}
	return memoryUUID, nil
}

func (r *SemanticService) Retrieve(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error) {
	rdbmsMemories, err := r.rdbmsMemoryRepo.FetchManyByConversationID(ctx, conversationID, lastK)
	if err != nil {
		return nil, fmt.Errorf("error fetching memories")
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
	vectorMemories, err := r.vectorMemoryRepo.Search(ctx, conversationID, query, topK)
	if err != nil {
		fmt.Printf("\n")
		return nil, fmt.Errorf("error searching for memories")
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
