package tool_test

import (
	"context"
	"testing"

	catalogdomain "github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
	catalogservice "github.com/mayankanup/commerce-ai-platform/internal/catalog/service"
	catalogtool "github.com/mayankanup/commerce-ai-platform/internal/catalog/tool"
)

type fakeCatalogRepository struct{}

func (r *fakeCatalogRepository) SearchProducts(
	ctx context.Context,
	request catalogdomain.SearchProductsRequest,
) (*catalogdomain.SearchProductsResponse, error) {

	return &catalogdomain.SearchProductsResponse{
		Products: []catalogdomain.ProductSearchResult{
			{
				SKU:   "NP-TSH-BLU-M",
				Name:  "Classic Cotton T-Shirt",
				Color: "Blue",
				Size:  "M",
				Price: 29.99,
			},
		},
	}, nil
}

func TestSearchProductsTool_Execute(t *testing.T) {

	// Arrange
	repository := &fakeCatalogRepository{}

	service := catalogservice.New(repository)

	tool := catalogtool.New(service)

	// Act
	result, err := tool.Execute(
		context.Background(),
		map[string]any{
			"query": "blue tshirt",
		},
	)

	// Assert

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response, ok := result.Content.(*catalogdomain.SearchProductsResponse)
	if !ok {
		t.Fatalf(
			"expected *SearchProductsResponse, got %T",
			result.Content,
		)
	}

	if len(response.Products) != 1 {
		t.Fatalf(
			"expected 1 product, got %d",
			len(response.Products),
		)
	}

	product := response.Products[0]

	if product.SKU != "NP-TSH-BLU-M" {
		t.Fatalf(
			"expected SKU NP-TSH-BLU-M, got %s",
			product.SKU,
		)
	}

	if product.Name != "Classic Cotton T-Shirt" {
		t.Fatalf(
			"expected product name Classic Cotton T-Shirt, got %s",
			product.Name,
		)
	}

	if product.Color != "Blue" {
		t.Fatalf(
			"expected color Blue, got %s",
			product.Color,
		)
	}

	if product.Size != "M" {
		t.Fatalf(
			"expected size M, got %s",
			product.Size,
		)
	}

	if product.Price != 29.99 {
		t.Fatalf(
			"expected price 29.99, got %f",
			product.Price,
		)
	}
}

func TestSearchProductsTool_MissingQuery(t *testing.T) {

	repository := &fakeCatalogRepository{}
	service := catalogservice.New(repository)
	tool := catalogtool.New(service)

	_, err := tool.Execute(
		context.Background(),
		map[string]any{},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSearchProductsTool_InvalidQuery(t *testing.T) {

	repository := &fakeCatalogRepository{}
	service := catalogservice.New(repository)
	tool := catalogtool.New(service)

	_, err := tool.Execute(
		context.Background(),
		map[string]any{
			"query": 123,
		},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
