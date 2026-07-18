package service

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
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

func (s *Service) CheckInventory(
	ctx context.Context,
	sku string,
) (*domain.CheckInventoryResult, error) {

	inventory, err := s.repository.GetInventoryBySKU(
		ctx,
		sku,
	)

	if err != nil {
		return nil, err
	}

	result := &domain.CheckInventoryResult{
		SKU: sku,
	}

	for _, item := range inventory {

		result.TotalAvailableQuantity += item.Available

		result.Warehouses = append(
			result.Warehouses,
			item,
		)
	}

	return result, nil
}
