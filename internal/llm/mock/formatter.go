package mock

import (
	"fmt"
	"strings"

	catalogdomain "github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
	inventorydomain "github.com/mayankanup/commerce-ai-platform/internal/inventory/domain"
)

func formatSearchProducts(
	response *catalogdomain.SearchProductsResponse,
) string {

	if len(response.Products) == 0 {
		return "I couldn't find any matching products."
	}

	var builder strings.Builder

	builder.WriteString("I found the following products:\n\n")

	for _, product := range response.Products {

		builder.WriteString(
			fmt.Sprintf(
				"- %s (%s, %s) - $%.2f\n",
				product.Name,
				product.Color,
				product.Size,
				product.Price,
			),
		)
	}

	return builder.String()
}

func formatInventory(
	result inventorydomain.CheckInventoryResult,
) string {

	if result.TotalQuantity == 0 {
		return fmt.Sprintf(
			"%s is currently out of stock.",
			result.SKU,
		)
	}

	var builder strings.Builder

	builder.WriteString(
		fmt.Sprintf(
			"%s is available.\n\n",
			result.SKU,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"Total Available: %d\n\n",
			result.TotalQuantity,
		),
	)

	builder.WriteString("Warehouse Inventory:\n")

	for _, warehouse := range result.Warehouses {

		builder.WriteString(
			fmt.Sprintf(
				"- %s (%s): %d available\n",
				warehouse.Warehouse.Name,
				warehouse.Warehouse.City,
				warehouse.Available,
			),
		)
	}

	return builder.String()
}
