package dto

type CreateBranchProductRequest struct {
	BranchID     string
	ProductID    string
	BuyPrice     float64
	SellingPrice float64
	SupplierID   string
}

type GetBranchProductFilter struct {
	Limit   int
	Offfset int
}

type GetBranchProductResponse struct {
	BranchID        string  `json:"branch_id"`
	BranchName      string  `json:"branch_name"`
	ProductID       string  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductCategory string  `json:"product_category"`
	Price           float64 `json:"price"`
	COGS            float64 `json:"cogs"`
}
