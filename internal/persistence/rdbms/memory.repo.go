package rdbms

import (
	"context"
	"database/sql"
	"memory/internal/persistence"
)

type MemoryRepo struct {
	db *sql.DB
}

func NewSqliteMemoryRepo(db *sql.DB) persistence.RdbmsMemoryRepo {
	return &MemoryRepo{db}
}

func (r *MemoryRepo) FetchOne(ctx context.Context, conversationID string) (persistence.RdbmsMemory, error) {
	return persistence.RdbmsMemory{}, nil
}

func (r *MemoryRepo) InsertOne(ctx context.Context, conversationID string, query string, response string) (persistence.RdbmsMemory, error) {
	return persistence.RdbmsMemory{}, nil
}
