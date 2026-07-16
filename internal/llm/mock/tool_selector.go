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

	if strings.Contains(lower, "shirt") ||
		strings.Contains(lower, "t-shirt") ||
		strings.Contains(lower, "tshirt") ||
		strings.Contains(lower, "hoodie") ||
		strings.Contains(lower, "jeans") {

		return &llm.ToolCall{
			Name: "search_products",

			Arguments: map[string]any{
				"query": prompt,
			},
		}, true
	}

	return nil, false
}
