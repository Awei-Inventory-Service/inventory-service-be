package inventory_transfer_item

import "github.com/inventory-service/resource/inventory_transfer_item"

func NewInventoryTransferItemDomain(
	inventoryTransferItemResource inventory_transfer_item.InventoryTransferItemResource,
) InventoryTransferItemDomain {
	return &inventoryTransferItemDomain{
		inventoryTransferItemResource: inventoryTransferItemResource,
	}
}
