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
		return uuid.UUID{}, fmt.Errorf("error creating conversation id")
	}
	createdAt := time.Now()
	_, err = r.conversationRepo.InsertOne(ctx, agent, user, conversationID, createdAt)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error creating conversation")
	}
	return conversationID, nil
}
