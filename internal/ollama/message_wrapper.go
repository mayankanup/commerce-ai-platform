package ollama

import "github.com/mayankanup/commerce-ai-platform/internal/llm"

func toOllamaMessage(
	message llm.Message,
) chatMessage {

	return chatMessage{
		Role:    string(message.Role),
		Content: message.Content,
	}
}

func fromOllamaMessage(
	message chatMessage,
) llm.Message {

	result := llm.Message{
		Role:    llm.Role(message.Role),
		Content: message.Content,
	}

	for _, tc := range message.ToolCalls {

		result.ToolCalls = append(
			result.ToolCalls,
			llm.ToolCall{
				Name:      tc.Function.Name,
				Arguments: tc.Function.Arguments,
			},
		)
	}

	return result
}
