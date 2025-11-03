package persistence

import (
	"github.com/google/uuid"
)

type RdbmsConversation struct {
	ID        string    `db:"id"`
	UUID      uuid.UUID `db:"uuid"`
	Agent     string    `db:"agent"`
	User      string    `db:"user"`
	CreatedAt string    `db:"created_at"`
}

type RdbmsMemory struct {
	ID             string    `db:"id"`
	UUID           uuid.UUID `db:"uuid"`
	ConversationID string    `db:"conversation_id"`
	Query          string    `db:"query"`
	Response       string    `db:"response"`
}

type VectorMemory struct {
	ID int
}

type CacheMemory struct {
	Query     string
	Response  string
	CreatedAt string
}
