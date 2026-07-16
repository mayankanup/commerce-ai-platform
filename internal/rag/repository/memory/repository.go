package memory

import (
	"context"
	"sort"
	"sync"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository"
)

type Repository struct {
	mu sync.RWMutex

	vectors map[string]repository.DocumentVector
}

func New() repository.Repository {
	return &Repository{
		vectors: make(map[string]repository.DocumentVector),
	}
}

func (r *Repository) Upsert(
	ctx context.Context,
	vectors []repository.DocumentVector,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, vector := range vectors {
		r.vectors[vector.ID] = vector
	}

	return nil
}

func (r *Repository) Search(
	ctx context.Context,
	queryEmbedding []float32,
	limit int,
) ([]repository.SearchResult, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	results := make([]repository.SearchResult, 0, len(r.vectors))

	for _, vector := range r.vectors {

		score := cosineSimilarity(
			queryEmbedding,
			vector.Embedding,
		)

		results = append(
			results,
			repository.SearchResult{
				DocumentVector: vector,
				Score:          score,
			},
		)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if limit > 0 && len(results) > limit {
		results = results[:limit]
	}

	return results, nil
}

func (r *Repository) Delete(
	ctx context.Context,
	documentID string,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	for id, vector := range r.vectors {

		if vector.DocumentID == documentID {
			delete(r.vectors, id)
		}
	}

	return nil
}

func (r *Repository) DeleteAll(
	ctx context.Context,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.vectors = make(map[string]repository.DocumentVector)

	return nil
}

func (r *Repository) Count(
	ctx context.Context,
) (int, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.vectors), nil
}
