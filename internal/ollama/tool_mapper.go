package ollama

import "github.com/mayankanup/commerce-ai-platform/internal/llm"

func toOllamaTools(
	definitions []llm.ToolDefinition,
) []toolDefinition {

	if len(definitions) == 0 {
		return nil
	}

	tools := make(
		[]toolDefinition,
		0,
		len(definitions),
	)

	for _, definition := range definitions {

		tools = append(
			tools,
			toolDefinition{
				Type: "function",
				Function: toolFunction{
					Name:        definition.Name,
					Description: definition.Description,
					Parameters:  definition.ParametersSchema,
				},
			},
		)
	}

	return tools
}
