package purchase

import "gorm.io/gorm"

func NewPurchaseResource(db *gorm.DB) PurchaseResource {
	return &purchaseResource{db: db}
}
