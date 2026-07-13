package sqlite

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
)

type InventoryRepository struct {
	db *Database
}

func NewInventoryRepository(db *Database) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (r *InventoryRepository) GetInventoryBySKU(
	ctx context.Context,
	sku string,
) ([]domain.WarehouseInventory, error) {

	const query = `
SELECT
    w.id,
    w.code,
    w.name,
    w.city,
    w.state,
    w.country,
    il.quantity_available,
    il.quantity_reserved,
    il.reorder_level
FROM product_variants pv
INNER JOIN inventory_levels il
    ON il.variant_id = pv.id
INNER JOIN warehouses w
    ON w.id = il.warehouse_id
WHERE pv.sku = ?
ORDER BY w.name;
`

	rows, err := r.db.DB().QueryContext(
		ctx,
		query,
		sku,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.WarehouseInventory

	for rows.Next() {

		var inventory domain.WarehouseInventory

		err := rows.Scan(
			&inventory.Warehouse.ID,
			&inventory.Warehouse.Code,
			&inventory.Warehouse.Name,
			&inventory.Warehouse.City,
			&inventory.Warehouse.State,
			&inventory.Warehouse.Country,
			&inventory.Available,
			&inventory.Reserved,
			&inventory.ReorderAt,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, inventory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
