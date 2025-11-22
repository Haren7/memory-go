package rdbms

import (
	"context"
	"database/sql"
	"memory/internal/persistence"
	"time"

	"github.com/google/uuid"
)

type ConversationRepo struct {
	db *sql.DB
}

func NewConversationRepo(db *sql.DB) persistence.ConversationRepoInterface {
	return &ConversationRepo{db}
}

func (r *ConversationRepo) FetchOne(ctx context.Context, conversationID uuid.UUID) (persistence.Conversation, error) {
	query := "SELECT id, uuid, agent, user, created_at FROM conversations WHERE uuid = $1"
	row := r.db.QueryRowContext(ctx, query, conversationID)
	var conversation persistence.Conversation
	err := row.Scan(&conversation.ID, &conversation.UUID, &conversation.Agent, &conversation.User, &conversation.CreatedAt)
	if err != nil {
		return persistence.Conversation{}, err
	}
	return conversation, nil
}

func (r *ConversationRepo) InsertOne(ctx context.Context, agent string, user string, conversationID uuid.UUID, createdAt time.Time) (int, error) {
	result, err := r.db.ExecContext(ctx, "INSERT INTO conversations (uuid, agent, user, created_at) VALUES ($1, $2, $3, $4)", conversationID, agent, user, createdAt)
	if err != nil {
		return 0, err
	}
	lastInsertId, _ := result.LastInsertId()
	return int(lastInsertId), nil
}
