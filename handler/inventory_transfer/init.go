package inventory_transfer

import "github.com/inventory-service/usecase/inventory_transfer"

func NewInventoryTransferHandler(inventoryTransferUsecase inventory_transfer.InventoryTransferUsecase) InventoryTransferHandler {
	return &inventoryTransferHandler{
		inventoryTransferUsecase: inventoryTransferUsecase,
	}
}