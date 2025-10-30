package product_snapshot

import "github.com/inventory-service/resource/product_snapshot"

func NewProductSnapshotDomain(
	productSnapshotResource product_snapshot.ProductSnaspshotResource,
) ProductSnapshotDomain {
	return &productSnapshotDomain{productSnapshotResource: productSnapshotResource}
}
