package domain

// DocumentChunk represents a searchable chunk of knowledge.
type DocumentChunk struct {
	ID string

	Document string

	Title string

	Content string

	Sequence int

	Score float64
}

// SearchRequest represents a semantic search request.
type SearchRequest struct {
	Query string

	TopK int
}

// SearchResponse contains matching chunks.
type SearchResponse struct {
	Chunks []DocumentChunk
}
