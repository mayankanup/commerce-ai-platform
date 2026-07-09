// package domain
package domain

// InventoryAvailability represents the inventory status
// of a purchasable SKU across all warehouses.
type InventoryAvailability struct {
	SKU         string
	ProductName string

	Warehouses []WarehouseInventory
}

// WarehouseInventory represents inventory
// for one warehouse.
type WarehouseInventory struct {
	Warehouse Warehouse

	Available int
	Reserved  int
	ReorderAt int
}

// Warehouse represents a fulfillment center.
type Warehouse struct {
	ID      int64
	Code    string
	Name    string
	City    string
	State   string
	Country string
}

type CheckInventoryResult struct {
	SKU string

	TotalQuantity int

	Warehouses []WarehouseInventory
}
