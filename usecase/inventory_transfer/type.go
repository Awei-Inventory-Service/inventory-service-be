package inventory_transfer

import (
	"context"

	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/inventory_transfer"
	"github.com/inventory-service/domain/inventory_transfer_item"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type InventoryTransferUsecase interface {
	Create(ctx context.Context, payload dto.CreateInventoryTransferRequest) (newData model.InventoryTransfer, errW *error_wrapper.ErrorWrapper)
	UpdateStatus(ctx context.Context, payload dto.UpdateInventoryTransferStatus) (errW *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, payload dto.GetListRequest) (result dto.GetInventoryTransferListResponse, errW *error_wrapper.ErrorWrapper)
}

type inventoryTransferUsecase struct {
	inventoryTransferItemDomain inventory_transfer_item.InventoryTransferItemDomain
	inventoryTransferDomain     inventory_transfer.InventoryTransferDomain
	inventoryDomain             inventory.InventoryDomain
	stockTransactionDomain      stocktransaction.StockTransactionDomain
	itemDomain                  item.ItemDomain
}
