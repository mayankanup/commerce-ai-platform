package service_test

import (
	"context"
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/domain"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/service"
)

type fakeRepository struct{}

func (r *fakeRepository) Search(
	ctx context.Context,
	request domain.SearchRequest,
) (*domain.SearchResponse, error) {

	return &domain.SearchResponse{
		Chunks: []domain.DocumentChunk{
			{
				ID:       "1",
				Document: "faq.md",
				Title:    "Return Policy",
				Content:  "Items can be returned within 30 days.",
				Score:    0.95,
			},
		},
	}, nil
}

func TestSearch(t *testing.T) {

	svc := service.New(
		&fakeRepository{},
	)

	response, err := svc.Search(
		context.Background(),
		domain.SearchRequest{
			Query: "return policy",
			TopK:  3,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(response.Chunks) != 1 {
		t.Fatalf(
			"expected 1 chunk, got %d",
			len(response.Chunks),
		)
	}

	if response.Chunks[0].Document != "faq.md" {
		t.Fatalf(
			"unexpected document: %s",
			response.Chunks[0].Document,
		)
	}
}
