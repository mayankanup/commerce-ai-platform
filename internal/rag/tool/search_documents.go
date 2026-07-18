package tool

import (
	"context"
	"fmt"
	"strings"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	"github.com/mayankanup/commerce-ai-platform/internal/llm"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/retriever"
)

const (
	defaultSearchLimit = 5
)

type SearchDocumentsTool struct {
	retriever *retriever.Retriever
}

func NewSearchDocumentsTool(
	retriever *retriever.Retriever,
) *SearchDocumentsTool {

	return &SearchDocumentsTool{
		retriever: retriever,
	}
}

func (t *SearchDocumentsTool) Definition() llm.ToolDefinition {

	return llm.ToolDefinition{
		Name: "search_documents",

		Description: "Searches company documents such as FAQs, return policy, shipping policy, store information and leadership information.",

		ParametersSchema: llm.JSONSchema{
			Type: "object",

			Properties: map[string]llm.Property{
				"query": {
					Type:        "string",
					Description: "Natural language search query.",
				},
			},

			Required: []string{
				"query",
			},
		},
	}
}

func (t *SearchDocumentsTool) Execute(
	ctx context.Context,
	input map[string]any,
) (*agent.ToolResult, error) {

	value, ok := input["query"]
	if !ok {
		return nil, fmt.Errorf("missing required parameter: query")
	}

	query, ok := value.(string)
	if !ok || strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("invalid parameter: query")
	}

	results, err := t.retriever.Search(
		ctx,
		query,
		defaultSearchLimit,
	)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {

		return &agent.ToolResult{
			Content: "No relevant documents found.",
		}, nil
	}

	var builder strings.Builder

	for i, result := range results {

		builder.WriteString(
			fmt.Sprintf(
				"Document: %s\n",
				result.DocumentVector.DocumentID,
			),
		)

		builder.WriteString(
			fmt.Sprintf(
				"Score: %.4f\n",
				result.Score,
			),
		)

		builder.WriteString(result.DocumentVector.Content)

		if i != len(results)-1 {
			builder.WriteString("\n\n---\n\n")
		}
	}

	return &agent.ToolResult{
		Content: builder.String(),
	}, nil
}
