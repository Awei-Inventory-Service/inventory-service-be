package product

import "go.mongodb.org/mongo-driver/mongo"

func NewProductRepository(mongoDb *mongo.Client, dbName string, collectionName string) ProductRepository {
	collection := mongoDb.Database(dbName).Collection(collectionName)

	return &productRepository{productCollection: collection}
}
