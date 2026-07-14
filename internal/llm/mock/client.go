package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
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

	//----------------------------------------------------
	// SECOND CALL (tool response)
	//----------------------------------------------------
	if last.Role == llm.ToolRole {

		if last.ToolName == "check_inventory" {

			var result domain.CheckInventoryResult

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
		}
	}

	//----------------------------------------------------
	// FIRST CALL (user prompt)
	//----------------------------------------------------

	prompt := strings.TrimSpace(messages[0].Content)

	for _, scenario := range c.scenarios {

		if strings.EqualFold(prompt, scenario.Prompt) {

			response := scenario.FirstResponse
			return &response, nil
		}
	}

	return &llm.Message{
		Role:    llm.AssistantRole,
		Content: "Sorry, I don't understand.",
	}, nil
}

func formatInventory(
	result domain.CheckInventoryResult,
) string {

	if result.TotalQuantity == 0 {
		return fmt.Sprintf(
			"%s is currently out of stock.",
			result.SKU,
		)
	}

	var builder strings.Builder

	builder.WriteString(
		fmt.Sprintf(
			"%s is available.\n\n",
			result.SKU,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"Total Available: %d\n\n",
			result.TotalQuantity,
		),
	)

	builder.WriteString("Warehouse Inventory:\n")

	for _, warehouse := range result.Warehouses {

		builder.WriteString(
			fmt.Sprintf(
				"- %s (%s): %d available\n",
				warehouse.Warehouse.Name,
				warehouse.Warehouse.City,
				warehouse.Available,
			),
		)
	}

	return builder.String()
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
