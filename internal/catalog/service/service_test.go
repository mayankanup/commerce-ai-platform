package service_test

import (
	"context"
	"testing"

	catalogdomain "github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
	catalogservice "github.com/mayankanup/commerce-ai-platform/internal/catalog/service"
)

type fakeRepository struct{}

func (r *fakeRepository) SearchProducts(
	ctx context.Context,
	request catalogdomain.SearchProductsRequest,
) (*catalogdomain.SearchProductsResponse, error) {

	return &catalogdomain.SearchProductsResponse{
		Products: []catalogdomain.ProductSearchResult{
			{
				SKU: "NP-TSH-BLU-M",
			},
		},
	}, nil
}

func TestDefaultLimit(t *testing.T) {

	service := catalogservice.New(&fakeRepository{})

	resp, err := service.SearchProducts(
		context.Background(),
		"blue",
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Products) != 1 {
		t.Fatal("expected one product")
	}
}
