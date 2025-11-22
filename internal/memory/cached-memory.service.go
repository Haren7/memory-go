package memory

import (
	"context"
	"fmt"
	"memory/internal/cache"
	"memory/internal/summarizer"
	"time"

	"github.com/google/uuid"
)

type CachedService struct {
	memoryRepo        cache.MemoryRepoInterface
	summarizerService summarizer.ServiceInterface
}

func NewCachedService(memoryRepo cache.MemoryRepoInterface, summarizerService summarizer.ServiceInterface) ServiceInterface {
	return &CachedService{
		memoryRepo:        memoryRepo,
		summarizerService: summarizerService,
	}
}

func (r *CachedService) Store(ctx context.Context, conversationID uuid.UUID, query, response string) (uuid.UUID, error) {
	memoryId, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error creating memory id")
	}
	createdAt := time.Now()
	summarizedResponse, err := r.summarizerService.Summarize(ctx, response)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error summarizing response")
	}
	r.memoryRepo.SetOne(ctx, conversationID, memoryId, query, summarizedResponse, createdAt)
	return memoryId, nil
}

func (r *CachedService) Retrieve(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error) {
	cacheMemories, err := r.memoryRepo.Get(ctx, conversationID, lastK)
	if err != nil {
		return nil, fmt.Errorf("error retrieving memories")
	}
	var memories []Memory
	for _, memory := range cacheMemories {
		memories = append(memories, Memory{
			ID:        memory.ID,
			Query:     memory.Query,
			Response:  memory.Response,
			CreatedAt: memory.CreatedAt,
		})
	}
	return memories, nil
}
