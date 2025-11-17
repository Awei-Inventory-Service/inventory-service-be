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

type GetProductionItem struct {
	SourceItemID    string  `json:"source_item_id"`
	SourceItemName  string  `json:"source_item_name"`
	InitialQuantity float64 `json:"initial_quantity"`
	Waste           float64 `json:"waste"`
	WastePercentage float64 `json:"waste_percentage"`
}

type GetProductionList struct {
	ProductionID   string              `json:"production_id"`
	FinalItemID    string              `json:"final_item_id"`
	FinalItemName  string              `json:"final_item_name"`
	FinalQuantity  float64             `json:"final_quantity"`
	FinalUnit      string              `json:"final_unit"`
	ProductionDate string              `json:"production_date"`
	BranchID       string              `json:"branch_id"`
	BranchName     string              `json:"branch_name"`
	SourceItems    []GetProductionItem `json:"source_items"`
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

type GetProductionFilter struct {
	ProductionID string `json:"production_id"`
	BranchID     string `json:"branch_id"`
	FinalItemID  string `json:"final_item_id"`
}

type DeleteProductionRequest struct {
	ProductionID string `json:"production_id"`
	BranchID     string `json:"branch_id"`
	UserID       string `json:"user_id"`
}

type UpdateProductionRequest struct {
	SourceItems    []SourceItemCreateProductionRequest `json:"source_items"`
	ProductionID   string                              `json:"production_id"`
	FinalItemID    string                              `json:"final_item_id"`
	FinalQuantity  float64                             `json:"final_quantity"`
	FinalUnit      string                              `json:"final_unit"`
	BranchID       string                              `json:"branch_id"`
	ProductionDate string                              `json:"production_date"`
	UserID         string                              `json:"user_id"`
}
