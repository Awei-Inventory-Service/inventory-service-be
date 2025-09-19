package purchase

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/domain/branch"
	branchitem "github.com/inventory-service/domain/branch_item"
	"github.com/inventory-service/domain/item"
	itempurchasechain "github.com/inventory-service/domain/item_purchase_chain"
	"github.com/inventory-service/domain/purchase"
	"github.com/inventory-service/domain/supplier"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	// "github.com/inventory-service/lib/error_wrapper"
)

// TODO: Change error with error_wrapper.ErrorWrapper
type PurchaseService interface {
	Create(c *gin.Context, payload dto.CreatePurchaseRequest) *error_wrapper.ErrorWrapper
	FindAll() ([]dto.GetPurchaseResponse, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, id, supplierId, branchId, itemId string, quantity float64, purchaseCost float64) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, id, userID string) *error_wrapper.ErrorWrapper
}

type purchaseService struct {
	purchaseDomain          purchase.PurchaseDomain
	supplierDomain          supplier.SupplierDomain
	branchDomain            branch.BranchDomain
	itemDomain              item.ItemDomain
	itemPurchaseChainDomain itempurchasechain.ItemPurchaseChainDomain
	branchItemDomain        branchitem.BranchItemDomain
}
