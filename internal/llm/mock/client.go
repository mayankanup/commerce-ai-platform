package mock

import (
	"context"
	"strings"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type Client struct {
	scenarios []Scenario
}

func New() *Client {

	return &Client{
		scenarios: defaultScenarios(),
	}
}

func (c *Client) Chat(
	ctx context.Context,
	messages []llm.Message,
	tools []llm.ToolDefinition,
) (*llm.Message, error) {

	_ = ctx
	_ = tools

	if len(messages) == 0 {
		return &llm.Message{
			Role: llm.AssistantRole,
		}, nil
	}

	last := messages[len(messages)-1]

	// Second call after tool execution.
	if last.Role == llm.ToolRole {

		for _, scenario := range c.scenarios {

			if last.ToolName == "check_inventory" {
				response := scenario.FinalResponse
				return &response, nil
			}
		}
	}

	prompt := strings.TrimSpace(messages[0].Content)

	for _, scenario := range c.scenarios {

		if strings.EqualFold(prompt, scenario.Prompt) {

			response := scenario.FirstResponse
			return &response, nil
		}
	}

	return &llm.Message{
		Role:    llm.AssistantRole,
		Content: "Sorry, I don't know how to answer that yet.",
	}, nil
}

func defaultScenarios() []Scenario {

	return []Scenario{
		{
			Prompt: "Do you have NP-TSH-BLK-M in stock?",

			FirstResponse: llm.Message{
				Role: llm.AssistantRole,
				ToolCalls: []llm.ToolCall{
					{
						Name: "check_inventory",
						Arguments: map[string]any{
							"sku": "NP-TSH-BLK-M",
						},
					},
				},
			},

			FinalResponse: llm.Message{
				Role:    llm.AssistantRole,
				Content: "Yes. NP-TSH-BLK-M is available in multiple warehouses.",
			},
		},
	}
}
