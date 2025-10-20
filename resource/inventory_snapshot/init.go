package inventory_snapshot

import (
	"github.com/inventory-service/constant"
	"github.com/inventory-service/resource/mongodb"
)

func NewInventorySnapshot(mongoDB mongodb.MongoDBClientWrapper) InventorySnapshotResource {
	collection := mongoDB.Database(constant.MONGO_DB_NAME).Collection(constant.MONGO_DB_INVENTORY_SNAPSHOT_COLLECTION)
	return &inventorySnapshotResource{inventorySnapshotCollection: collection}
}
