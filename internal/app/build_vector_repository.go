package app

import (
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	ragrepo "github.com/mayankanup/commerce-ai-platform/internal/rag/repository"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository/memory"
)

func buildVectorRepository(
	cfg *config.Config,
) (ragrepo.Repository, error) {

	switch cfg.RAG.VectorDB.Provider {

	case "memory":

		return memory.New(), nil

	case "chroma":

		return nil, fmt.Errorf(
			"chroma repository not implemented yet",
		)

	default:

		return nil, fmt.Errorf(
			"unsupported vector repository provider: %s",
			cfg.RAG.VectorDB.Provider,
		)
	}
}
