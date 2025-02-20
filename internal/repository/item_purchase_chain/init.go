package itempurchasechain

import (
	"github.com/inventory-service/internal/repository/mongodb"
)

func NewItemPurchaseChainRepository(mongoDb mongodb.MongoDBClientWrapper, dbName string, collectionName string) ItemPurchaseChainRepository {
	collection := mongoDb.Database(dbName).Collection(collectionName)
	return &itemPurchaseChainRepository{itemPurchaseChainCollection: collection}
}
