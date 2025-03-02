package dto

type CreateSalesRequest struct {
	BranchID  string
	ProductID string
	Type      string
	Quantity  int
}
