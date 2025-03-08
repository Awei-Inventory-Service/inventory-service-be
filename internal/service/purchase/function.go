package purchase

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (p *purchaseService) Create(c *gin.Context, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper {
	var (
		errChan = make(chan *error_wrapper.ErrorWrapper, 3) // Buffered channel untuk 3 goroutines
		wg      sync.WaitGroup
		errors  []*error_wrapper.ErrorWrapper // Slice untuk menyimpan semua error
	)

	// supplier check
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := p.supplierRepository.FindByID(supplierId)
		errChan <- err
	}()

	// branch check
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := p.branchRepository.FindByID(branchId)
		errChan <- err
	}()

	// item check
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := p.itemRepository.FindByID(itemId)
		errChan <- err
	}()

	// Goroutine untuk menutup channel setelah semua selesai
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Kumpulkan semua error dari channel
	for err := range errChan {
		if err != nil {
			errors = append(errors, err) // Simpan error
		}
	}

	// Jika ada error, return error pertama yang ditemukan
	if len(errors) > 0 {
		return errors[0]
	}

	// Lanjutkan proses pembuatan purchase
	newPurchase, errW := p.purchaseRepository.Create(supplierId, branchId, itemId, quantity, purchaseCost)
	if errW != nil {
		return errW
	}
	fmt.Println("iNI NEW PURCHASE", newPurchase)
	fmt.Println("Creating item purchase chain")
	errW = p.itemPurchaseChainRepository.Create(c, itemId, branchId, *newPurchase)

	if errW != nil {
		return errW
	}
	fmt.Println("Done creating ipc")
	return nil
}

func (p *purchaseService) FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	purchases, err := p.purchaseRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return purchases, nil
}

func (p *purchaseService) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase, err := p.purchaseRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (p *purchaseService) Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper {
	errChan := make(chan *error_wrapper.ErrorWrapper, 3)

	// Supplier check
	go func() {
		_, err := p.supplierRepository.FindByID(supplierId)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Branch check
	go func() {
		_, err := p.branchRepository.FindByID(branchId)
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// Item check
	go func() {
		_, err := p.itemRepository.FindByID(itemId)
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

	err := p.purchaseRepository.Update(id, supplierId, branchId, itemId, quantity, purchaseCost)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := p.purchaseRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
