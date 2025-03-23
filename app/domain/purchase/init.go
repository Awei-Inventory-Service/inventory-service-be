package purchase

import "github.com/inventory-service/app/resource/purchase"

func NewPurchaseDomain(purchaseResource purchase.PurchaseResource) PurchaseDomain {
	return &purchaseDomain{purchaseResource: purchaseResource}
}
