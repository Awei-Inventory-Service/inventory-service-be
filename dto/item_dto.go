package dto

type CreateItemRequest struct {
	Name             string                         `json:"name" binding:"required"`
	SupplierID       string                         `json:"supplier_id"`
	Category         string                         `json:"category" binding:"required"`
	Price            float64                        `json:"price"`
	Unit             string                         `json:"unit" binding:"required"`
	PortionSize      float64                        `json:"portion_size"`
	ItemCompositions []CreateItemCompositionRequest `json:"item_compositions"`
}

type CreateItemCompositionRequest struct {
	ItemID string  `json:"item_id" binding:"required"`
	Ratio  float64 `json:"ratio" binding:"required"`
	Notes  string  `json:"notes"`
}

type UpdateItemCompositionRequest struct {
	ItemID string  `json:"item_id" binding:"required"`
	Ratio  float64 `json:"ratio" binding:"required"`
	Notes  string  `json:"notes"`
}

type UpdateItemRequest struct {
	Name             string                         `json:"name" binding:"required"`
	SupplierID       string                         `json:"supplier_id"`
	Category         string                         `json:"category" binding:"required"`
	Price            float64                        `json:"price"`
	Unit             string                         `json:"unit" binding:"required"`
	PortionSize      float64                        `json:"portion_size"`
	ItemCompositions []UpdateItemCompositionRequest `json:"item_compositions"`
}
