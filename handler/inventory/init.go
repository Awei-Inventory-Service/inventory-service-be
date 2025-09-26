package inventory

import "github.com/inventory-service/usecase/inventory"

func NewInventoryHandler(inventoryUsecase inventory.InventoryUsecase) InventoryHandler {
	return &inventoryHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
