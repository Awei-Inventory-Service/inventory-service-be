package model

type Status string

const (
	StatusInUse   Status = "in-use"
	StatusNotUsed Status = "not-used"
	StatusUsed    Status = "used"
)

type ItemPurchaseChain struct {
	ItemID       string   `json:"item_id"`
	BranchID     string   `json:"branch_id"`
	Purchase     Purchase `json:"purchase"`
	Quantity     int      `json:"quantity"`
	Status       Status   `json:"status"`
	SalesRecords []Sales  `json:"sales"`
}

type ItemPurchaseChainGet struct {
	ID           string   `json:"_id"`
	ItemID       string   `json:"item_id"`
	BranchID     string   `json:"branch_id"`
	Purchase     Purchase `json:"purchase"`
	Quantity     int      `json:"quantity"`
	Status       Status   `json:"status"`
	SalesRecords []Sales  `json:"sales"`
}
