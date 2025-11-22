package memory

import (
	"time"

	"github.com/google/uuid"
)

type SortOrder string

const (
	ASC  SortOrder = "asc"
	DESC SortOrder = "desc"
)

type SortKey string

const (
	CREATED_AT SortKey = "created_at"
)

type RerankerOpts struct {
	SortOrder SortOrder
	SortKey   SortKey
}

type Memory struct {
	ID        uuid.UUID
	Query     string
	Response  string
	CreatedAt time.Time
}
