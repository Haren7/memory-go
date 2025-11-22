package conversation

import (
	"context"
	"fmt"
	"memory/internal/persistence"
	"time"

	"github.com/google/uuid"
)

type ConversationServiceInterface interface {
	Create(ctx context.Context, agent, user string) (uuid.UUID, error)
	Exists(ctx context.Context, conversationID uuid.UUID) (bool, error)
}

type ConversationService struct {
	conversationRepo persistence.ConversationRepoInterface
}

func NewConversationService(conversationRepo persistence.ConversationRepoInterface) ConversationServiceInterface {
	return &ConversationService{
		conversationRepo: conversationRepo,
	}
}

func (r *ConversationService) Create(ctx context.Context, agent, user string) (uuid.UUID, error) {
	conversationID, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("conversation: error creating conversation id, %w", err)
	}
	createdAt := time.Now()
	_, err = r.conversationRepo.InsertOne(ctx, agent, user, conversationID, createdAt)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("conversation: error creating conversation, %w", err)
	}
	return conversationID, nil
}

func (r *ConversationService) Exists(ctx context.Context, conversationID uuid.UUID) (bool, error) {
	_, err := r.conversationRepo.FetchOne(ctx, conversationID)
	if err != nil {
		return false, fmt.Errorf("conversation: error checking if conversation exists, %w", err)
	}
	return true, nil
}
