package product

import (
	"github.com/inventory-service/app/resource/mongodb"
)

func NewProductResource(mongoDb mongodb.MongoDBClientWrapper, dbName string, collectionName string) ProductResource {
	collection := mongoDb.Database(dbName).Collection(collectionName)

	return &productResource{productCollection: collection}
}
