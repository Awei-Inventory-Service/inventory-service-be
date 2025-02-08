package purchase

import (
	"errors"

	"github.com/inventory-service/internal/model"
)

func (p *purchaseService) Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) error {
	errChan := make(chan error, 3)
	// supplier check
	go func() {
		_, err := p.supplierRepository.FindByID(supplierId)
		if err != nil {
			errChan <- errors.New("supplier not found")
		} else {
			errChan <- nil
		}
	}()

	// branch check
	go func() {
		_, err := p.branchRepository.FindByID(branchId)
		if err != nil {
			errChan <- errors.New("supplier not found")
		} else {
			errChan <- nil
		}
	}()

	// item check
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

	err := p.purchaseRepository.Create(supplierId, branchId, itemId, quantity, purchaseCost)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseService) FindAll() ([]model.Purchase, error) {
	purchases, err := p.purchaseRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return purchases, nil
}

func (p *purchaseService) FindByID(id string) (*model.Purchase, error) {
	purchase, err := p.purchaseRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (p *purchaseService) Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) error {
	errChan := make(chan error, 3)

	// Supplier check
	go func() {
		_, err := p.supplierRepository.FindByID(supplierId)
		if err != nil {
			errChan <- errors.New("supplier not found")
		} else {
			errChan <- nil
		}
	}()

	// Branch check
	go func() {
		_, err := p.branchRepository.FindByID(branchId)
		if err != nil {
			errChan <- errors.New("branch not found")
		} else {
			errChan <- nil
		}
	}()

	// Item check
	go func() {
		_, err := p.itemRepository.FindByID(itemId)
		if err != nil {
			errChan <- errors.New("item not found")
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

func (p *purchaseService) Delete(id string) error {
	err := p.purchaseRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
