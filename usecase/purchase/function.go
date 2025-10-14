package purchase

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/constant"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseService) Create(c *gin.Context, payload dto.CreatePurchaseRequest) *error_wrapper.ErrorWrapper {
	var (
		errChan = make(chan *error_wrapper.ErrorWrapper, 3) // Buffered channel for 3 goroutines
		wg      sync.WaitGroup
	)

	// Supplier check
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := p.supplierDomain.FindByID(payload.SupplierID); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Branch check
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := p.branchDomain.FindByID(payload.BranchID); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Item check
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := p.itemDomain.FindByID(c, payload.ItemID); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(errChan) // Close the channel when all goroutines are done
	}()

	// Collect all errors from the channel
	for err := range errChan {
		if err != nil {
			return err // If there's an error, return immediately
		}
	}

	userId := c.GetHeader("user_id")
	// All validation completed, domain handles all inventory logic
	purchase, errW := p.purchaseDomain.Create(payload, userId)

	if errW != nil {
		return errW
	}
	referenceType := string(constant.Purchasing)

	errW = p.stockTransactionDomain.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.ItemID,
		Type:                "IN",
		IssuerID:            userId,
		Quantity:            payload.Quantity,
		Cost:                payload.PurchaseCost,
		Unit:                payload.Unit,
		Reference:           purchase.UUID,
		ReferenceType:       &referenceType,
	})

	_, _, errW = p.inventoryDomain.SyncBranchItem(c, payload.BranchID, purchase.ItemID)

	return errW
}

func (p *purchaseService) FindAll() ([]dto.GetPurchaseResponse, *error_wrapper.ErrorWrapper) {
	purchases, err := p.purchaseDomain.FindAll()
	if err != nil {
		return nil, err
	}

	return purchases, nil
}

func (p *purchaseService) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase, err := p.purchaseDomain.FindByID(id)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (p *purchaseService) Update(ctx context.Context, id string, payload dto.UpdatePurchaseRequest) *error_wrapper.ErrorWrapper {
	errChan := make(chan *error_wrapper.ErrorWrapper, 3)

	// Supplier check
	go func() {
		_, err := p.supplierDomain.FindByID(payload.SupplierID)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Branch check
	go func() {
		_, err := p.branchDomain.FindByID(payload.BranchID)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Item check
	go func() {
		_, err := p.itemDomain.FindByID(ctx, payload.ItemID)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	for i := 0; i < 3; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	// 1. Update purchase
	errW := p.purchaseDomain.Update(id, payload)
	if errW != nil {
		return errW
	}
	// 2. Invalidate old stock transaction and create new one
	errW = p.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": id,
		},
	}, payload.UserID)

	if errW != nil {
		return errW
	}

	// 3. Add new stock transaction
	referenceType := string(constant.Purchasing)

	errW = p.stockTransactionDomain.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.ItemID,
		Type:                "IN",
		IssuerID:            payload.UserID,
		Quantity:            payload.Quantity,
		Cost:                payload.PurchaseCost,
		Unit:                payload.Unit,
		ReferenceType:       &referenceType,
		Reference:           id,
	})

	if errW != nil {
		return errW
	}

	// 4. Sync branch item

	_, _, errW = p.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)

	return errW
}

func (p *purchaseService) Delete(ctx context.Context, id, userID string) *error_wrapper.ErrorWrapper {
	// Domain handles all inventory logic including sync
	deletedPurchase, errW := p.purchaseDomain.Delete(ctx, id, userID)

	if errW != nil {
		return errW
	}

	// 2. Invalidate stock transaction
	errW = p.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": id,
		},
	}, userID)

	if errW != nil {
		return errW
	}

	_, _, errW = p.inventoryDomain.SyncBranchItem(ctx, deletedPurchase.BranchID, deletedPurchase.ItemID)

	return errW
}
