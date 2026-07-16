package indexer_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/embedding/mock"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/chunker"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/indexer"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/loader"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository/memory"
)

func TestIndexer_Index(t *testing.T) {

	documentsDir := filepath.Join(
		"..",
		"..",
		"..",
		"..",
		"docs",
	)

	loader := loader.NewFilesystemLoader(
		documentsDir,
	)

	chunker := chunker.New()

	embeddingProvider := mock.New()

	vectorRepository := memory.New()

	indexer := indexer.New(
		loader,
		chunker,
		embeddingProvider,
		vectorRepository,
	)

	err := indexer.Index(
		context.Background(),
	)
	if err != nil {
		t.Fatalf("index failed: %v", err)
	}

	count, err := indexer.Count(
		context.Background(),
	)
	if err != nil {
		t.Fatalf("count failed: %v", err)
	}

	if count == 0 {
		t.Fatal("expected vectors to be indexed")
	}
}

func TestIndexer_Rebuild(t *testing.T) {

	documentsDir := filepath.Join(
		"..",
		"..",
		"..",
		"..",
		"docs",
	)

	loader := loader.NewFilesystemLoader(
		documentsDir,
	)

	chunker := chunker.New()

	embeddingProvider := mock.New()

	vectorRepository := memory.New()

	indexer := indexer.New(
		loader,
		chunker,
		embeddingProvider,
		vectorRepository,
	)

	err := indexer.Rebuild(
		context.Background(),
	)
	if err != nil {
		t.Fatalf("rebuild failed: %v", err)
	}

	count, err := indexer.Count(
		context.Background(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if count == 0 {
		t.Fatal("expected vectors after rebuild")
	}
}
