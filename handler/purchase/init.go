package purchase

import "github.com/inventory-service/usecase/purchase"

func NewPurchaseController(purchaseService purchase.PurchaseService) PurchaseController {
	return &purchaseController{
		purchaseService: purchaseService,
	}
}
