package itempurchasechain

import (
	"github.com/inventory-service/resource/mongodb"
)

func NewItemPurchaseChainResource(mongoDb mongodb.MongoDBClientWrapper, dbName string, collectionName string) ItemPurchaseChainResource {
	collection := mongoDb.Database(dbName).Collection(collectionName)
	return &itemPurchaseChainResource{itemPurchaseChainCollection: collection}
}
