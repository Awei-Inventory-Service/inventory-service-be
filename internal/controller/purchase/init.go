package purchase

import "github.com/inventory-service/internal/service/purchase"

func NewPurchaseController(purchaseService purchase.PurchaseService) PurchaseController {
	return &purchaseController{
		purchaseService: purchaseService,
	}
}
