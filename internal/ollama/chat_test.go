package ollama_test

import (
	"context"
	"testing"
	"time"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
	"github.com/mayankanup/commerce-ai-platform/internal/ollama"
)

func TestChat(t *testing.T) {

	client := ollama.New(
		ollama.Options{
			Endpoint: "http://localhost:11434",
			Model:    "qwen3.6:latest",
			Timeout:  120 * time.Second,
		},
	)

	response, err := client.Chat(
		context.Background(),
		[]llm.Message{
			{
				Role:    "user",
				Content: "Reply with exactly the word hello",
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if response.Content == "" {
		t.Fatal("empty response")
	}
}
