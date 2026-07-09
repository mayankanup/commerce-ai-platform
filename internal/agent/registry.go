package agent

import (
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type Registry struct {
	tools map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

func (r *Registry) Register(tool Tool) {
	r.tools[tool.Definition().Name] = tool
}

func (r *Registry) Get(name string) (Tool, error) {

	tool, ok := r.tools[name]

	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrToolNotFound, name)
	}

	return tool, nil
}

func (r *Registry) Definitions() []llm.ToolDefinition {

	result := make([]llm.ToolDefinition, 0, len(r.tools))

	for _, tool := range r.tools {
		result = append(
			result,
			tool.Definition(),
		)
	}

	return result
}
