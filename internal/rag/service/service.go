package service

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/domain"
)

type Service struct {
	repository domain.Repository
}

func New(
	repository domain.Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) Search(
	ctx context.Context,
	request domain.SearchRequest,
) (*domain.SearchResponse, error) {

	return s.repository.Search(
		ctx,
		request,
	)
}
