package domain

import "context"

type Repository interface {
	GetInventoryBySKU(
		ctx context.Context,
		sku string,
	) ([]WarehouseInventory, error)
}
