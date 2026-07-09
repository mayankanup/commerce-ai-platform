package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
)

var (
	ErrSKUNotFound = errors.New("sku not found")
)

type SQLiteRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewSQLiteRepository(
	db *sql.DB,
	logger *slog.Logger,
) *SQLiteRepository {

	return &SQLiteRepository{
		db:     db,
		logger: logger,
	}
}

const getInventoryBySKUQuery = `
SELECT
    pv.sku,
    p.name,

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

JOIN products p
    ON p.id = pv.product_id

JOIN inventory_levels il
    ON il.variant_id = pv.id

JOIN warehouses w
    ON w.id = il.warehouse_id

WHERE pv.sku = ?

ORDER BY w.name;
`

func (r *SQLiteRepository) GetInventoryBySKU(
	ctx context.Context,
	sku string,
) (*domain.InventoryAvailability, error) {

	rows, err := r.db.QueryContext(
		ctx,
		getInventoryBySKUQuery,
		sku,
	)
	if err != nil {
		return nil, fmt.Errorf("query inventory: %w", err)
	}
	defer rows.Close()

	inventory := &domain.InventoryAvailability{}

	found := false

	for rows.Next() {

		found = true

		var warehouse domain.Warehouse
		var warehouseInventory domain.WarehouseInventory

		err := rows.Scan(
			&inventory.SKU,
			&inventory.ProductName,

			&warehouse.ID,
			&warehouse.Code,
			&warehouse.Name,
			&warehouse.City,
			&warehouse.State,
			&warehouse.Country,

			&warehouseInventory.Available,
			&warehouseInventory.Reserved,
			&warehouseInventory.ReorderAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan inventory: %w", err)
		}

		warehouseInventory.Warehouse = warehouse

		inventory.Warehouses = append(
			inventory.Warehouses,
			warehouseInventory,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate inventory: %w", err)
	}

	if !found {
		return nil, ErrSKUNotFound
	}

	r.logger.Info(
		"Inventory lookup completed",
		"sku", sku,
		"warehouseCount", len(inventory.Warehouses),
	)

	return inventory, nil
}
