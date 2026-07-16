package mock

import (
	"context"
	"encoding/json"
	"strings"

	catalogdomain "github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
	inventorydomain "github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Chat(
	ctx context.Context,
	messages []llm.Message,
	tools []llm.ToolDefinition,
) (*llm.Message, error) {

	_ = ctx

	if len(messages) == 0 {
		return &llm.Message{
			Role:    llm.AssistantRole,
			Content: "No messages received.",
		}, nil
	}

	last := messages[len(messages)-1]

	//--------------------------------------------------------
	// Second LLM call
	// Tool has already executed.
	//--------------------------------------------------------

	if last.Role == llm.ToolRole {

		switch last.ToolName {

		case "check_inventory":

			var result inventorydomain.CheckInventoryResult

			if err := json.Unmarshal(
				[]byte(last.Content),
				&result,
			); err != nil {
				return nil, err
			}

			return &llm.Message{
				Role:    llm.AssistantRole,
				Content: formatInventory(result),
			}, nil

		case "search_products":

			var response catalogdomain.SearchProductsResponse

			if err := json.Unmarshal([]byte(last.Content), &response); err != nil {
				return nil, err
			}

			return &llm.Message{
				Role:    llm.AssistantRole,
				Content: formatSearchProducts(&response),
			}, nil
		}

		return &llm.Message{
			Role:    llm.AssistantRole,
			Content: "Tool executed successfully.",
		}, nil
	}

	//--------------------------------------------------------
	// First LLM call
	// User prompt
	//--------------------------------------------------------

	prompt := strings.TrimSpace(last.Content)

	toolCall, ok := selectTool(
		prompt,
		tools,
	)

	if ok {

		return &llm.Message{
			Role: llm.AssistantRole,
			ToolCalls: []llm.ToolCall{
				*toolCall,
			},
		}, nil
	}

	return &llm.Message{
		Role:    llm.AssistantRole,
		Content: "Sorry, I couldn't determine which tool to use.",
	}, nil
}
