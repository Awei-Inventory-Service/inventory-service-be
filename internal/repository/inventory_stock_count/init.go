package inventorystockcount

import "go.mongodb.org/mongo-driver/mongo"

func NewInventoryStockCountRepository(mongoDb *mongo.Client, dbName string, collectionName string) InventoryStockCountRepository {
	collection := mongoDb.Database(dbName).Collection(collectionName)

	return &inventoryStockCountRepository{inventoryStockCountCollection: collection}
}
