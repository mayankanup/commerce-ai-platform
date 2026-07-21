package prompt

import "github.com/mayankanup/commerce-ai-platform/internal/llm"

type Builder struct {
	systemPrompt string
}

func New(systemPrompt string) *Builder {
	return &Builder{
		systemPrompt: systemPrompt,
	}
}

func (b *Builder) Build(
	userMessage string,
) []llm.Message {

	return []llm.Message{
		{
			Role:    llm.SystemRole,
			Content: b.systemPrompt,
		},
		{
			Role:    llm.UserRole,
			Content: userMessage,
		},
	}
}
