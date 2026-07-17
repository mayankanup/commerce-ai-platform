package ollama

type EmbeddingRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type EmbeddingResponse struct {
	Embedding []float32 `json:"embedding"`
}

type chatRequest struct {
	Model string `json:"model"`

	Messages []chatMessage `json:"messages"`

	Stream bool `json:"stream"`
}

type chatMessage struct {
	Role string `json:"role"`

	Content string `json:"content"`
}

type chatResponse struct {
	Model string `json:"model"`

	Message chatMessage `json:"message"`

	Done bool `json:"done"`
}
