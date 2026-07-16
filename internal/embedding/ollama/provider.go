package ollama

import (
	"context"
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/embedding"
	ollamaclient "github.com/mayankanup/commerce-ai-platform/internal/ollama"
)

const (
	defaultDimensions = 384
)

type Provider struct {
	client *ollamaclient.Client

	model string
}

func New(
	client *ollamaclient.Client,
	model string,
) embedding.Provider {

	return &Provider{
		client: client,
		model:  model,
	}
}

func (p *Provider) Embed(
	ctx context.Context,
	text string,
) ([]float32, error) {

	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	return p.client.Embedding(
		ctx,
		ollamaclient.EmbeddingRequest{
			Model:  p.model,
			Prompt: text,
		},
	)
}

func (p *Provider) Model() string {
	return p.model
}

func (p *Provider) Dimensions() int {
	return defaultDimensions
}
