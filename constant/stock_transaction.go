package constant

type ReferenceType string

const (
	Production        ReferenceType = "production"
	Purchasing        ReferenceType = "purchasing"
	InventoryTransfer ReferenceType = "inventory-transfer"
	Sales             ReferenceType = "sales"
)
