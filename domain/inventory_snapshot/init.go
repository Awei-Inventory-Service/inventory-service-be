package inventory_snapshot

import "github.com/inventory-service/resource/inventory_snapshot"

func NewInventorySnapshotDomain(inventorySnapshoutResource inventory_snapshot.InventorySnapshotResource) InventorySnapshotDomain {
	return &inventorySnapshotDomain{
		inventorySnapshotResource: inventorySnapshoutResource,
	}
}
