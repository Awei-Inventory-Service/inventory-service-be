package purchase

import "github.com/inventory-service/app/usecase/purchase"

func NewPurchaseController(purchaseService purchase.PurchaseService) PurchaseController {
	return &purchaseController{
		purchaseService: purchaseService,
	}
}
