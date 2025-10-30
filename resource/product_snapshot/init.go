package product_snapshot

import (
	"github.com/inventory-service/constant"
	"github.com/inventory-service/resource/mongodb"
)

func NewProductSnapshotResource(
	mongoDB mongodb.MongoDBClientWrapper,
) ProductSnaspshotResource {
	collection := mongoDB.Database(constant.MONGO_DB_NAME).Collection(constant.MONGO_DB_INVENTORY_SNAPSHOT_COLLECTION)
	return &productSnapshotResource{productSnapshotCollection: collection}
}
