package domain

type SearchProductsRequest struct {
	Query string
	Limit int
}

type ProductSearchResult struct {
	SKU  string
	Name string

	Color string
	Size  string

	Price float64
}

type SearchProductsResponse struct {
	Products []ProductSearchResult
}
