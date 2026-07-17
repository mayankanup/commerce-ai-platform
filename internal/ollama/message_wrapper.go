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

	return llm.Message{
		Role:    llm.Role(message.Role),
		Content: message.Content,
	}
}
