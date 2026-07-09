package llm

// Message represents one message exchanged with the LLM.
type Message struct {
	Role Role

	// Natural language content.
	Content string

	// Tool calls requested by the model.
	ToolCalls []ToolCall

	// Tool response messages use this field.
	ToolName string
}
