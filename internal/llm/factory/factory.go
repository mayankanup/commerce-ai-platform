package factory

import (
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
	"github.com/mayankanup/commerce-ai-platform/internal/llm/mock"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
)

func NewClient(cfg *config.Config) (llm.Client, error) {

	switch llm.Provider(cfg.LLM.Provider) {

	case llm.ProviderMock:
		return mock.New(), nil

	case llm.ProviderOllama:
		return nil, fmt.Errorf("ollama provider not implemented")

	default:
		return nil, fmt.Errorf(
			"unsupported llm provider: %s",
			cfg.LLM.Provider,
		)
	}
}
