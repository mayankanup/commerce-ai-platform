package llm

import "context"

// Client represents any Large Language Model provider.
type Client interface {
	Chat(
		ctx context.Context,
		messages []Message,
		tools []ToolDefinition,
	) (*Message, error)
}
