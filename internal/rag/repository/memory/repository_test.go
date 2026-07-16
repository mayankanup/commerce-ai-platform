package memory_test

import (
	"context"
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository/memory"
)

func TestUpsertAndCount(t *testing.T) {

	repo := memory.New()

	err := repo.Upsert(
		context.Background(),
		[]repository.DocumentVector{
			{
				ID:         "1",
				DocumentID: "company",
				ChunkID:    "1",
				Content:    "Northwind Outfitters",
				Embedding:  []float32{1, 0},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	count, err := repo.Count(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if count != 1 {
		t.Fatalf("expected 1 vector, got %d", count)
	}
}

func TestSearch(t *testing.T) {

	repo := memory.New()

	repo.Upsert(
		context.Background(),
		[]repository.DocumentVector{
			{
				ID:         "1",
				DocumentID: "faq",
				ChunkID:    "1",
				Content:    "Return policy",
				Embedding:  []float32{1, 0},
			},
			{
				ID:         "2",
				DocumentID: "stores",
				ChunkID:    "1",
				Content:    "Store locations",
				Embedding:  []float32{0, 1},
			},
		},
	)

	results, err := repo.Search(
		context.Background(),
		[]float32{1, 0},
		1,
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(results) != 1 {
		t.Fatal("expected one result")
	}

	if results[0].DocumentID != "faq" {
		t.Fatalf(
			"expected faq, got %s",
			results[0].DocumentID,
		)
	}
}

func TestDelete(t *testing.T) {

	repo := memory.New()

	repo.Upsert(
		context.Background(),
		[]repository.DocumentVector{
			{
				ID:         "1",
				DocumentID: "company",
				ChunkID:    "1",
				Embedding:  []float32{1, 0},
			},
		},
	)

	err := repo.Delete(
		context.Background(),
		"company",
	)
	if err != nil {
		t.Fatal(err)
	}

	count, _ := repo.Count(context.Background())

	if count != 0 {
		t.Fatal("vector should have been deleted")
	}
}

func TestDeleteAll(t *testing.T) {

	repo := memory.New()

	repo.Upsert(
		context.Background(),
		[]repository.DocumentVector{
			{
				ID:         "1",
				DocumentID: "a",
				Embedding:  []float32{1},
			},
			{
				ID:         "2",
				DocumentID: "b",
				Embedding:  []float32{2},
			},
		},
	)

	repo.DeleteAll(context.Background())

	count, _ := repo.Count(context.Background())

	if count != 0 {
		t.Fatal("expected repository to be empty")
	}
}
