package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

func (c *Client) Chat(
	ctx context.Context,
	messages []llm.Message,
	tools []llm.ToolDefinition,
) (*llm.Message, error) {

	_ = tools // AI-002-01 ignores tools

	requestMessages := make(
		[]chatMessage,
		0,
		len(messages),
	)

	for _, message := range messages {
		requestMessages = append(
			requestMessages,
			toOllamaMessage(message),
		)
	}

	request := chatRequest{
		Model: c.model,

		Messages: requestMessages,

		Stream: false,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.endpoint+"/api/chat",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf(
			"ollama returned status %d",
			resp.StatusCode,
		)
	}

	var response chatResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	message := fromOllamaMessage(response.Message)

	return &message, nil
}
