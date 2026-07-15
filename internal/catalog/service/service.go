package service

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
	"github.com/mayankanup/commerce-ai-platform/internal/catalog/repository"
)

type Service struct {
	repository repository.Repository
}

func New(
	repository repository.Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) SearchProducts(
	ctx context.Context,
	query string,
) (*domain.SearchProductsResponse, error) {

	return s.repository.SearchProducts(
		ctx,
		domain.SearchProductsRequest{
			Query: query,
			Limit: 5,
		},
	)
}
