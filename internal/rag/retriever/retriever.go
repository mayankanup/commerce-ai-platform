package retriever

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/embedding"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository"
)

const (
	DefaultTopK = 5
)

type Retriever struct {
	embedding embedding.Provider

	repository repository.Repository
}

func New(
	embedding embedding.Provider,
	repository repository.Repository,
) *Retriever {

	return &Retriever{
		embedding:  embedding,
		repository: repository,
	}
}

func (r *Retriever) Search(
	ctx context.Context,
	query string,
	limit int,
) ([]repository.SearchResult, error) {

	if limit <= 0 {
		limit = DefaultTopK
	}

	queryEmbedding, err := r.embedding.Embed(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}

	return r.repository.Search(
		ctx,
		queryEmbedding,
		limit,
	)
}

func (r *Retriever) SearchDefault(
	ctx context.Context,
	query string,
) ([]repository.SearchResult, error) {

	return r.Search(
		ctx,
		query,
		DefaultTopK,
	)
}
