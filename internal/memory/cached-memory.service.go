package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/haren7/minimal-memory/internal/cache"
	"github.com/haren7/minimal-memory/internal/summarizer"

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
		return uuid.UUID{}, fmt.Errorf("cached: error creating memory id, %w", err)
	}
	createdAt := time.Now()
	summarizedResponse, err := r.summarizerService.Summarize(ctx, response)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cached: error summarizing response, %w", err)
	}
	r.memoryRepo.SetOne(ctx, conversationID, memoryId, query, summarizedResponse, createdAt)
	return memoryId, nil
}

func (r *CachedService) Retrieve(ctx context.Context, conversationID uuid.UUID, lastK int) ([]Memory, error) {
	cacheMemories, err := r.memoryRepo.Get(ctx, conversationID, lastK)
	if err != nil {
		return nil, fmt.Errorf("cached: error retrieving memories, %w", err)
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
