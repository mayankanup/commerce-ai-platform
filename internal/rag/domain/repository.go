package domain

import "context"

// Repository provides semantic search capabilities.
type Repository interface {
	Search(
		ctx context.Context,
		request SearchRequest,
	) (*SearchResponse, error)
}
