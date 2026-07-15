package sqlite

import (
	"context"

	catalogdomain "github.com/mayankanup/commerce-ai-platform/internal/catalog/domain"
)

type SQLiteCatalogRepository struct {
	db *Database
}

func NewCatalogRepository(
	db *Database,
) *SQLiteCatalogRepository {

	return &SQLiteCatalogRepository{
		db: db,
	}
}

func (r *SQLiteCatalogRepository) SearchProducts(
	ctx context.Context,
	request catalogdomain.SearchProductsRequest,
) (*catalogdomain.SearchProductsResponse, error) {

	query := `
SELECT
    pv.sku,
    p.name,
    pv.color,
    pv.size,
    pv.price
FROM product_variants pv
JOIN products p
ON p.id=pv.product_id
WHERE
    LOWER(p.name) LIKE LOWER(?)
 OR LOWER(pv.color) LIKE LOWER(?)
LIMIT ?`

	search := "%" + request.Query + "%"

	rows, err := r.db.DB().QueryContext(
		ctx,
		query,
		search,
		search,
		request.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &catalogdomain.SearchProductsResponse{}

	for rows.Next() {

		var item catalogdomain.ProductSearchResult

		err := rows.Scan(
			&item.SKU,
			&item.Name,
			&item.Color,
			&item.Size,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}

		response.Products = append(
			response.Products,
			item,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}
