package app

import (
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
	"github.com/mayankanup/commerce-ai-platform/internal/llm/mock"
	"github.com/mayankanup/commerce-ai-platform/internal/ollama"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
)

func buildLLMClient(
	cfg *config.Config,
) (llm.Client, error) {

	switch llm.Provider(cfg.LLM.Provider) {

	case llm.ProviderMock:

		return mock.New(), nil

	case llm.ProviderOllama:

		return ollama.New(
			ollama.Options{
				Endpoint: cfg.LLM.Ollama.Endpoint,
				Model:    cfg.LLM.Ollama.Model,
				Timeout:  cfg.LLM.Ollama.Timeout,
			},
		), nil

	default:

		return nil, fmt.Errorf(
			"unsupported llm provider: %s",
			cfg.LLM.Provider,
		)
	}
}
