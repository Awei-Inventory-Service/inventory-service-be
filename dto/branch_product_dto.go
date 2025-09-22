package dto

type CreateBranchProductRequest struct {
	BranchID     string
	ProductID    string
	BuyPrice     float64
	SellingPrice float64
	SupplierID   string
}
