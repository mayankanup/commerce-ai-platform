package repository

import "context"

// Repository stores and searches document vectors.
type Repository interface {

	// Upsert inserts or updates vectors.
	Upsert(
		ctx context.Context,
		vectors []DocumentVector,
	) error

	// Search returns the nearest vectors.
	Search(
		ctx context.Context,
		queryEmbedding []float32,
		limit int,
	) ([]SearchResult, error)

	// Delete removes vectors belonging to a document.
	Delete(
		ctx context.Context,
		documentID string,
	) error

	// DeleteAll removes every vector.
	DeleteAll(
		ctx context.Context,
	) error

	// Count returns the number of indexed vectors.
	Count(
		ctx context.Context,
	) (int, error)
}
