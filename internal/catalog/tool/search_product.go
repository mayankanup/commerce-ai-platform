package tool

import (
	"context"
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	"github.com/mayankanup/commerce-ai-platform/internal/catalog/service"
	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type SearchProductsTool struct {
	service *service.Service
}

func New(
	service *service.Service,
) *SearchProductsTool {

	return &SearchProductsTool{
		service: service,
	}
}

func (t *SearchProductsTool) Definition() llm.ToolDefinition {

	return llm.ToolDefinition{
		Name: "search_products",

		Description: "Searches the product catalog using natural language like blue t-shirt, hoodie, jeans or shoes.",

		ParametersSchema: llm.JSONSchema{
			Type: "object",

			Properties: map[string]llm.Property{
				"query": {
					Type:        "string",
					Description: "Natural language product search query",
				},
			},

			Required: []string{
				"query",
			},
		},
	}
}

func (t *SearchProductsTool) Execute(
	ctx context.Context,
	input map[string]any,
) (*agent.ToolResult, error) {

	value, ok := input["query"]
	if !ok {
		return nil, fmt.Errorf("missing required parameter: query")
	}

	query, ok := value.(string)
	if !ok || query == "" {
		return nil, fmt.Errorf("invalid parameter: query")
	}

	result, err := t.service.SearchProducts(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}

	return &agent.ToolResult{
		Content: result,
	}, nil
}
