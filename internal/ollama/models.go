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

	Tools []toolDefinition `json:"tools,omitempty"`

	Stream bool `json:"stream"`
}

type chatMessage struct {
	Role string `json:"role"`

	Content string `json:"content"`

	ToolCalls []toolCall `json:"tool_calls,omitempty"`
}

type chatResponse struct {
	Model string `json:"model"`

	Message chatMessage `json:"message"`

	Done bool `json:"done"`
}

type toolDefinition struct {
	Type string `json:"type"`

	Function toolFunction `json:"function"`
}

type toolFunction struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Parameters any `json:"parameters"`
}

type toolCall struct {
	Function toolFunctionCall `json:"function"`
}

type toolFunctionCall struct {
	Name string `json:"name"`

	Arguments map[string]any `json:"arguments"`
}
