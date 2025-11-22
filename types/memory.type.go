package types

import "time"

// Semantic Memory
type SemanticMemory struct {
	ID        string
	Query     string
	Response  string
	CreatedAt time.Time
}

type StoreSemanticMemoryInput struct {
	ConversationID string
	Query          string
	Response       string
}

type StoreSemanticMemoryOutput struct {
	MemoryID string
}

type RetrieveSemanticMemoryInput struct {
	ConversationID string
	Query          string
	TopK           int
}

type RetrieveSemanticMemoryOutput struct {
	Memories        []Memory
	SimilarMemories []SemanticMemory
}

// Short Term Memory
type Memory struct {
	ID        string
	Query     string
	Response  string
	CreatedAt time.Time
}

type StoreShortTermMemoryInput struct {
	Query          string
	Response       string
	ConversationID string
}

type StoreShortTermMemoryOutput struct {
	MemoryID string
}

type RetrieveShortTermMemoryInput struct {
	TopK           int
	ConversationID string
}

type RetrieveShortTermMemoryOutput struct {
	Memories []Memory
}
