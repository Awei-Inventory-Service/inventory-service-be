package dto

type CreateInvoiceRequest struct {
	FileURL		string `json:"file_url" binding:"required"`
	Amount		float64 `json:"amount" binding:"required"`
	AmountOwed	float64 `json:"amount_owed" binding:"required"`
	Notes		string `json:"notes" binding:"required"`
	InvoiceDate string `json:"invoice_date" binding:"required"`
}

type UpdateInvoiceRequest struct {
	FileURL		string `json:"file_url" binding:"required"`
	Amount		float64 `json:"amount" binding:"required"`
	AmountOwed	float64 `json:"amount_owed" binding:"required"`
	Notes		string `json:"notes" binding:"required"`
	InvoiceDate string `json:"invoice_date" binding:"required"`
}