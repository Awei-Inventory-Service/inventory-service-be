package inventorystockcount

import (
	"github.com/inventory-service/app/resource/mongodb"
)

func NewInventoryStockCountResource(mongoDb mongodb.MongoDBClientWrapper, dbName string, collectionName string) InventoryStockCountResource {
	collection := mongoDb.Database(dbName).Collection(collectionName)

	return &inventoryStockCountResource{inventoryStockCountCollection: collection}
}
