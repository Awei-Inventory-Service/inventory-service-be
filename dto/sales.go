package dto

type CreateSalesData struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required"`
	Type      string  `json:"type" binding:"required"`
}

type CreateSalesRequest struct {
	BranchID        string            `json:"branch_id" binding:"required"`
	SalesData       []CreateSalesData `json:"sales_data" binding:"required"`
	TransactionDate string            `json:"transaction_date"`
}

type GetSalesResponse struct {
	BranchID    string  `json:"branch_id"`
	BranchName  string  `json:"branch_name"`
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    float64 `json:"quantity"`
}

type SalesGroupedByDateResponse struct {
	TransactionDate string             `json:"transaction_date"`
	TotalSales      int                `json:"total_sales"`
	TotalRevenue    float64            `json:"total_revenue"`
	TotalProfit     float64            `json:"total_profit"`
	Sales           []GetSalesResponse `json:"sales"`
}

type BranchSalesData struct {
	BranchID     string             `json:"branch_id"`
	BranchName   string             `json:"branch_name"`
	TotalSales   int                `json:"total_sales"`
	TotalRevenue float64            `json:"total_revenue"`
	TotalProfit  float64            `json:"total_profit"`
	Sales        []GetSalesResponse `json:"sales"`
}

type SalesGroupedByDateAndBranchResponse struct {
	TransactionDate string             `json:"transaction_date"`
	TotalSales      int                `json:"total_sales"`
	TotalRevenue    float64            `json:"total_revenue"`
	TotalProfit     float64            `json:"total_profit"`
	BranchID        string             `json:"branch_id"`
	BranchName      string             `json:"branch_name"`
	Sales           []GetSalesResponse `json:"sales"`
}
