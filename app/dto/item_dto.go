package dto

type CreateItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	SupplierID  string  `json:"supplier_id" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Unit        string  `json:"unit" binding:"required"`
	PortionSize float64 `json:"portion_size"`
}

type UpdateItemRequest struct {
	Name       string  `json:"name" binding:"required"`
	Category   string  `json:"category" binding:"required"`
	SupplierID string  `json:"supplier_id" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Unit       string  `json:"unit" binding:"required"`
}
