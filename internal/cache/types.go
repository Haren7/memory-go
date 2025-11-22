package cache

import (
	"time"

	"github.com/google/uuid"
)

type Memory struct {
	ID        uuid.UUID
	Query     string
	Response  string
	CreatedAt time.Time
}
