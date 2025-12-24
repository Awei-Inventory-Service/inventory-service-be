package dto

type CreateSalesData struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required"`
	Type      string  `json:"type" binding:"required"`
	Cost      float64
	Price     float64
}

type UpdateSalesRequest struct {
	BranchID        string            `json:"branch_id"`
	TransactionDate string            `json:"transaction_date"`
	SalesID         string            `json:"sales_id"`
	SalesData       []UpdateSalesData `json:"sales_data"`
	Cost            float64
	Price           float64
	UserID          string
}

type UpdateSalesData struct {
	ProductID string  `json:"product_id"`
	Quantity  float64 `json:"quantity"`
	Type      string  `json:"type"`
}

type CreateSalesRequest struct {
	BranchID        string            `json:"branch_id" binding:"required"`
	SalesData       []CreateSalesData `json:"sales_data" binding:"required"`
	TransactionDate string            `json:"transaction_date"`
}

type GetSalesListResponse struct {
	SalesID         string                    `json:"sales_id"`
	TransactionDate string                    `json:"transaction_date"`
	BranchID        string                    `json:"branch_id"`
	BranchName      string                    `json:"branch_name"`
	SalesProducts   []GetSalesProductResponse `json:"sales_products"`
}

type GetSalesProductResponse struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    float64 `json:"quantity"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Cost        float64 `json:"cost"`
}
