package model

import "time"

type ProductType string

const (
	ProductTypeConsignment ProductType = "consignment"
	ProductTypeProduced    ProductType = "produced"
)

type Product struct {
	UUID         string      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Code         string      `gorm:"type:varchar(255);not null" json:"code"`
	Name         string      `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Category     string      `gorm:"type:varchar(255);not null" json:"category" validate:"required"`
	Type         ProductType `gorm:"type:varchar(20);not null"`
	Unit         string      `gorm:"type:varchar(255);not null" json:"unit" validate:"required"`
	SellingPrice float64     `gorm:"type:decimal(10,2);not null" json:"selling_price" validate:"required,gt=0"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`

	ProductComposition []ProductRecipe `gorm:"foreignKey:ProductID" json:"product_recipe,omitempty"`
}

type ProductRecipe struct {
	UUID      string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	ProductID string    `gorm:"type:uuid;not null;index" json:"product_id" validate:"required"` // Fixed: string not uint
	ItemID    string    `gorm:"type:uuid;not null;index" json:"item_id" validate:"required"`    // Fixed: string not uint
	Amount    float64   `gorm:"type:decimal(10,4);not null" json:"amount" validate:"required,gt=0"`
	Notes     string    `gorm:"type:text" json:"notes"`
	Unit      string    `gorm:"type:text" json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product Product `gorm:"foreignKey:ProductID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product,omitempty"`
	Item    Item    `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"item,omitempty"`
}

func (Product) TableName() string {
	return "products"
}

func (ProductRecipe) TableName() string {
	return "product_recipes"
}
