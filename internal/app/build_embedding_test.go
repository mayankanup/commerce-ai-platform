package app

import (
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
)

func TestBuildMockEmbeddingProvider(
	t *testing.T,
) {

	cfg := &config.Config{}

	cfg.Embedding.Provider = "mock"

	provider, err := buildEmbeddingProvider(
		cfg,
	)
	if err != nil {
		t.Fatal(err)
	}

	if provider.Model() != "mock" {
		t.Fatal("unexpected model")
	}
}

func TestBuildOllamaEmbeddingProvider(
	t *testing.T,
) {

	cfg := &config.Config{}

	cfg.Embedding.Provider = "ollama"
	cfg.Embedding.Model = "all-minilm:latest"
	cfg.Embedding.Endpoint = "http://localhost:11434"

	provider, err := buildEmbeddingProvider(
		cfg,
	)
	if err != nil {
		t.Fatal(err)
	}

	if provider.Model() != "all-minilm:latest" {
		t.Fatal("unexpected model")
	}
}
