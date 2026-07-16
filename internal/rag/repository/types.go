package repository

// DocumentVector represents one indexed chunk stored
// in the vector repository.
type DocumentVector struct {
	ID string

	DocumentID string
	ChunkID    string

	Content string

	Embedding []float32

	Metadata map[string]string
}

// SearchResult represents one search hit.
type SearchResult struct {
	DocumentVector

	Score float64
}
