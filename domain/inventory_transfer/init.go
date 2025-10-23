package inventory_transfer

import (
	"github.com/inventory-service/resource/inventory_transfer"
	"github.com/inventory-service/resource/inventory_transfer_item"
)

func NewInventoryTransferDomain(
	inventoryTransferResource inventory_transfer.InventoryTransferResource,
	inventoryTransferItemResource inventory_transfer_item.InventoryTransferItemResource,
) InventoryTransferDomain {
	return &inventoryTransferDomain{
		inventoryTransferResource:     inventoryTransferResource,
		inventoryTransferItemResource: inventoryTransferItemResource,
	}
}
