package repository

import (
	"context"

	catalogdomain "github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
)

type Repository interface {
	SearchProducts(
		ctx context.Context,
		request catalogdomain.SearchProductsRequest,
	) (*catalogdomain.SearchProductsResponse, error)
}
