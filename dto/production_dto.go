package dto

type SourceItemCreateProductionRequest struct {
	SourceItemID    string  `json:"source_item_id" binding:"required"`
	InitialQuantity float64 `json:"initial_quantity" binding:"required"`
	InitialUnit     string  `json:"initial_unit" binding:"required"`
}

type CreateProductionRequest struct {
	SourceItems    []SourceItemCreateProductionRequest `json:"source_items" binding:"required"`
	FinalItemID    string                              `json:"final_item_id" binding:"required"`
	FinalQuantity  float64                             `json:"final_quantity" binding:"required"`
	FinalUnit      string                              `json:"final_unit" binding:"required"`
	BranchID       string                              `json:"branch_id" binding:"required"`
	ProductionDate string                              `json:"production_date" binding:"required"`
	UserID         string                              `json:"user_id"`
}

type GetProduction struct {
	UUID            string  `json:"uuid"`
	SourceItemID    string  `json:"source_item_id"`
	SourceItemName  string  `json:"source_item_name"`
	FinalItemID     string  `json:"final_item_id"`
	FinalItemName   string  `json:"final_item_name"`
	Waste           float64 `json:"waste"`
	WastePercentage float64 `json:"waste_percentage"`
	BranchID        string  `json:"branch_id"`
	BranchName      string  `json:"branch_name"`
	ProductionDate  string  `json:"production_date"`
}
