package purchase

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type PurchaseDomain interface {
	Create(payload dto.CreatePurchaseRequest, userId string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	FindAll() ([]dto.GetPurchaseResponse, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id string, payload dto.UpdatePurchaseRequest) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, id, userID string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (payload []dto.GetPurchaseResponse, errW *error_wrapper.ErrorWrapper)
}

type purchaseDomain struct {
	purchaseResource         purchase.PurchaseResource
	inventoryResource        inventory.InventoryResource
	stockTransactionResource stocktransaction.StockTransactionResource
	itemResource             item.ItemResource
}
