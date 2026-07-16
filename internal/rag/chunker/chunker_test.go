package chunker_test

import (
	"strings"
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/chunker"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/domain"
)

func TestChunkDocument(t *testing.T) {

	text := strings.Repeat(
		"Northwind Outfitters sells quality clothing. ",
		50,
	)

	document := domain.Document{
		Name:    "company.md",
		Content: text,
	}

	c := chunker.NewWithOptions(
		200,
		50,
	)

	chunks := c.Chunk(document)

	if len(chunks) < 2 {
		t.Fatal("expected multiple chunks")
	}

	if chunks[0].Document != "company.md" {
		t.Fatal("incorrect document")
	}

	if chunks[0].Sequence != 0 {
		t.Fatal("first sequence should be zero")
	}

	for _, chunk := range chunks {

		if chunk.Content == "" {
			t.Fatal("chunk should not be empty")
		}
	}
}

func TestEmptyDocument(t *testing.T) {

	c := chunker.New()

	chunks := c.Chunk(
		domain.Document{},
	)

	if len(chunks) != 0 {
		t.Fatal("expected zero chunks")
	}
}
