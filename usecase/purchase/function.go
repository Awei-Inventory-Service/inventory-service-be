package purchase

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
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
	_, errW := p.purchaseDomain.Create(payload, userId)
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

func (p *purchaseService) Update(ctx context.Context, id, supplierId, branchId, itemId string, quantity float64, purchaseCost float64) *error_wrapper.ErrorWrapper {
	errChan := make(chan *error_wrapper.ErrorWrapper, 3)

	// Supplier check
	go func() {
		_, err := p.supplierDomain.FindByID(supplierId)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Branch check
	go func() {
		_, err := p.branchDomain.FindByID(branchId)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Item check
	go func() {
		_, err := p.itemDomain.FindByID(ctx, itemId)
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

	err := p.purchaseDomain.Update(id, supplierId, branchId, itemId, quantity, purchaseCost)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseService) Delete(ctx context.Context, id, userID string) *error_wrapper.ErrorWrapper {
	// Domain handles all inventory logic including sync
	_, err := p.purchaseDomain.Delete(ctx, id, userID)
	return err
}
