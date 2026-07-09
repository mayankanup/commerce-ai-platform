package tool

import (
	"context"
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	"github.com/mayankanup/commerce-ai-platform/internal/inventory/service"
	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type CheckInventoryTool struct {
	service *service.Service
}

func New(
	service *service.Service,
) *CheckInventoryTool {

	return &CheckInventoryTool{
		service: service,
	}
}

func (t *CheckInventoryTool) Definition() llm.ToolDefinition {

	return llm.ToolDefinition{
		Name: "check_inventory",

		Description: "Checks inventory availability for a SKU.",

		ParametersSchema: llm.JSONSchema{
			Type: "object",

			Properties: map[string]llm.Property{
				"sku": {
					Type:        "string",
					Description: "Product SKU",
				},
			},

			Required: []string{
				"sku",
			},
		},
	}
}

func (t *CheckInventoryTool) Execute(
	ctx context.Context,
	input map[string]any,
) (*agent.ToolResult, error) {

	value, ok := input["sku"]
	if !ok {
		return nil, fmt.Errorf("missing required parameter: sku")
	}

	sku, ok := value.(string)
	if !ok || sku == "" {
		return nil, fmt.Errorf("invalid parameter: sku")
	}

	result, err := t.service.CheckInventory(
		ctx,
		sku,
	)
	if err != nil {
		return nil, err
	}

	return &agent.ToolResult{
		Content: result,
	}, nil
}
