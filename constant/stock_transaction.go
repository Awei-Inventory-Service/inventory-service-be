package constant

const (
	Production         = "production"
	Purchasing         = "purchasing"
	InventoryTransfer  = "inventory-transfer"
	Sales              = "sales"
	DeleteAction       = "delete-action"
	ReversalProduction = "reversal-production"
)

var (
	ReferenceTypeMap = map[string]string{
		Production: ReversalProduction,
	}
)
