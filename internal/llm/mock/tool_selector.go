package mock

import (
	"regexp"
	"strings"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

var skuRegex = regexp.MustCompile(`[A-Z]{2}-[A-Z]{3}-[A-Z]{3}-[A-Z]`)

func selectTool(
	prompt string,
	tools []llm.ToolDefinition,
) (*llm.ToolCall, bool) {

	lower := strings.ToLower(prompt)

	if strings.Contains(lower, "stock") ||
		strings.Contains(lower, "available") ||
		strings.Contains(lower, "inventory") {

		sku := skuRegex.FindString(prompt)

		return &llm.ToolCall{
			Name: "check_inventory",
			Arguments: map[string]any{
				"sku": sku,
			},
		}, true
	}

	return nil, false
}
