package product

import (
	"github.com/inventory-service/internal/repository/mongodb"
)

func NewProductRepository(mongoDb mongodb.MongoDBClientWrapper, dbName string, collectionName string) ProductRepository {
	collection := mongoDb.Database(dbName).Collection(collectionName)

	return &productRepository{productCollection: collection}
}
