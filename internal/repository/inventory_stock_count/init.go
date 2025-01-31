package inventorystockcount

import (
	"github.com/inventory-service/internal/repository/mongodb"
)

func NewInventoryStockCountRepository(mongoDb mongodb.MongoDBClient, dbName string, collectionName string) InventoryStockCountRepository {
	collection := mongoDb.Database(dbName).Collection(collectionName)

	return &inventoryStockCountRepository{inventoryStockCountCollection: collection}
}
