package mock

import "github.com/mayankanup/commerce-ai-platform/internal/llm"

// Scenario represents one deterministic conversation
// supported by the Mock LLM.
type Scenario struct {
	FirstResponse llm.Message

	FinalResponse llm.Message
}
