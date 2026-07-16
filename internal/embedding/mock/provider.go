package mock

import (
	"context"
	"hash/fnv"

	"github.com/mayankanup/commerce-ai-platform/internal/embedding"
)

const (
	vectorDimensions = 384
)

type Provider struct{}

func New() embedding.Provider {
	return &Provider{}
}

func (p *Provider) Embed(
	ctx context.Context,
	text string,
) ([]float32, error) {

	vector := make([]float32, vectorDimensions)

	hash := fnv.New64a()
	hash.Write([]byte(text))

	value := hash.Sum64()

	for i := range vector {
		vector[i] = float32((value>>uint(i%32))&0xFF) / 255.0
	}

	return vector, nil
}

func (p *Provider) Model() string {
	return "mock"
}

func (p *Provider) Dimensions() int {
	return vectorDimensions
}
