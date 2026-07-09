package agent

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type Tool interface {

	// Definition describes the tool to the LLM.
	Definition() llm.ToolDefinition

	// Execute executes the tool.
	Execute(
		ctx context.Context,
		input map[string]any,
	) (*ToolResult, error)
}
