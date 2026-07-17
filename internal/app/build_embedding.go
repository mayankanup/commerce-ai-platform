package app

import (
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/embedding"
	embeddingmock "github.com/mayankanup/commerce-ai-platform/internal/embedding/mock"
	embeddingollama "github.com/mayankanup/commerce-ai-platform/internal/embedding/ollama"
	ollamaclient "github.com/mayankanup/commerce-ai-platform/internal/ollama"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
)

func buildEmbeddingProvider(
	cfg *config.Config,
) (embedding.Provider, error) {

	switch cfg.Embedding.Provider {

	case "mock":

		return embeddingmock.New(), nil

	case "ollama":

		client := ollamaclient.New(
			ollamaclient.Options{
				Endpoint: cfg.Embedding.Endpoint,
			},
		)

		return embeddingollama.New(
			client,
			cfg.Embedding.Model,
		), nil

	default:

		return nil, fmt.Errorf(
			"unsupported embedding provider: %s",
			cfg.Embedding.Provider,
		)
	}
}
