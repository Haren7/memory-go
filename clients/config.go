package clients

type ShortTermMemoryClientConfig struct {
}

type SemanticMemoryClientConfig struct {
	ContextWindowSize int
	OpenAIApiKey      string
}
