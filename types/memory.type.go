package types

import "time"

// Semantic Memory
type SemanticMemory struct {
	Query     string
	Response  string
	CreatedAt time.Time
}

type StoreSemanticMemoryInput struct {
	ConversationID string
	Query          string
	Response       string
}

type RetrieveSemanticMemoryInput struct {
	ConversationID string
	Query          string
	TopK           int
}

type RetrieveSemanticMemoryOutput struct {
	Memories []SemanticMemory
}

// Short Term Memory
type ShortTermMemory struct {
	Query     string
	Response  string
	CreatedAt time.Time
}

type StoreShortTermMemoryInput struct {
	Query    string
	Response string
}

type RetrieveShortTermMemoryInput struct {
	ConversationID string
	TopK           int
}

type RetrieveShortTermMemoryOutput struct {
	Memories []ShortTermMemory
}

// Episodic Memory
