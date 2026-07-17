package ollama_test

import (
	"context"
	"testing"

	embeddingollama "github.com/mayankanup/commerce-ai-platform/internal/embedding/ollama"
	ollamaclient "github.com/mayankanup/commerce-ai-platform/internal/ollama"
)

func TestEmbed(t *testing.T) {

	client := ollamaclient.New(
		ollamaclient.Options{
			Endpoint: "http://localhost:11434",
		},
	)

	provider := embeddingollama.New(
		client,
		"all-minilm:latest",
	)

	vector, err := provider.Embed(
		context.Background(),
		"Northwind Outfitters return policy",
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(vector) != provider.Dimensions() {
		t.Fatalf(
			"expected %d dimensions, got %d",
			provider.Dimensions(),
			len(vector),
		)
	}
}

func TestModel(t *testing.T) {

	ollamaClientOptions := ollamaclient.Options{
		Endpoint: "http://localhost:11434",
	}
	client := ollamaclient.New(
		ollamaClientOptions,
	)

	provider := embeddingollama.New(
		client,
		"all-minilm:latest",
	)

	if provider.Model() != "all-minilm:latest" {
		t.Fatal("unexpected model")
	}
}
