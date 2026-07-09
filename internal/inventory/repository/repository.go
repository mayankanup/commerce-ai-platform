package repository

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
)

type Repository interface {

	// GetInventoryBySKU returns inventory availability
	// across all warehouses for a SKU.
	GetInventoryBySKU(
		ctx context.Context,
		sku string,
	) (*domain.InventoryAvailability, error)
}
