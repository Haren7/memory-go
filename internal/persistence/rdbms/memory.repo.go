package rdbms

import (
	"context"
	"database/sql"
	"fmt"
	"memory/internal/persistence"
	"time"

	"github.com/google/uuid"
)

type MemoryRepo struct {
	db        *sql.DB
	tableName string
}

func NewMemoryRepo(db *sql.DB) persistence.MemoryRepoInterface {
	return &MemoryRepo{db: db, tableName: "memories"}
}

func NewFaissMemoryRepo(db *sql.DB) persistence.MemoryRepoInterface {
	return &MemoryRepo{db: db, tableName: "memories_meta"}
}

func (r *MemoryRepo) FetchOne(ctx context.Context, conversationID uuid.UUID) (persistence.Memory, error) {
	query := fmt.Sprintf(`SELECT id, uuid, conversation_id, query, response, created_at FROM %s WHERE conversation_id = $1`, r.tableName)
	row := r.db.QueryRowContext(ctx, query, conversationID)
	var memory persistence.Memory
	err := row.Scan(&memory.ID, &memory.UUID, &memory.ConversationID, &memory.Query, &memory.Response, &memory.CreatedAt)
	if err != nil {
		return persistence.Memory{}, err
	}
	return memory, nil
}

func (r *MemoryRepo) FetchMany(ctx context.Context, memoryIds []int) ([]persistence.Memory, error) {
	query := fmt.Sprintf(`SELECT id, uuid, conversation_id, query, response, created_at FROM %s WHERE id IN ($1)`, r.tableName)
	rows, err := r.db.QueryContext(ctx, query, memoryIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var memories []persistence.Memory
	for rows.Next() {
		var memory persistence.Memory
		err := rows.Scan(&memory.ID, &memory.UUID, &memory.ConversationID, &memory.Query, &memory.Response, &memory.CreatedAt)
		if err != nil {
			return nil, err
		}
		memories = append(memories, memory)
	}
	return memories, nil
}

func (r *MemoryRepo) FetchManyByConversationID(ctx context.Context, conversationID uuid.UUID, limit int) ([]persistence.Memory, error) {
	query := fmt.Sprintf(`SELECT id, uuid, conversation_id, query, response, created_at FROM %s WHERE conversation_id = $1`, r.tableName)
	rows, err := r.db.QueryContext(ctx, query, conversationID)
	if err != nil {
		return nil, err
	}
	var memories []persistence.Memory
	for rows.Next() {
		var memory persistence.Memory
		err = rows.Scan(&memory.ID, &memory.UUID, &memory.ConversationID, &memory.Query, &memory.Response, &memory.CreatedAt)
		if err != nil {
			return nil, err
		}
		memories = append(memories, memory)
	}
	return memories, nil
}

func (r *MemoryRepo) InsertOne(ctx context.Context, conversationID uuid.UUID, memoryID uuid.UUID, query string, response string, createdAt time.Time) (int, error) {
	result, err := r.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO %s (conversation_id, memory_id, query, response, created_at) VALUES ($1, $2, $3, $4, $5)`, r.tableName), conversationID, memoryID, query, response, createdAt)
	if err != nil {
		return 0, err
	}
	lastInsertId, _ := result.LastInsertId()
	return int(lastInsertId), nil
}
