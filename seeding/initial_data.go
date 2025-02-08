package seeding

import (
	"context"
	"fmt"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/mongodb"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(pgDB *gorm.DB) {
	mongodDb, err := mongodb.InitMongoDB()

	if err != nil {
		panic(err)
	}
	var (
		createdUserIDs     []string
		createdSupplierIDS []string
		createdItem        []model.Item
	)
	for i := 0; i < 10; i++ {

		hashed, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user := model.User{
			Username: fmt.Sprintf("username%d", i),
			Email:    fmt.Sprintf("user%d@gmail.com", i),
			Password: string(hashed),
			Role:     model.RoleBranchManager,
		}

		result := pgDB.Create(&user)

		if result.Error != nil {
			continue
		}
		createdUserIDs = append(createdUserIDs, user.UUID)
		fmt.Println("iNI USER", user)
	}

	if len(createdUserIDs) < 5 {
		panic("There is no enough user for creating branches")
	}

	for i := 0; i < 5; i++ {
		branch := model.Branch{
			Name:            fmt.Sprintf("Branch %d", i),
			Location:        fmt.Sprintf("Jalan ke-%d", i),
			BranchManagerID: createdUserIDs[i],
		}

		result := pgDB.Create(&branch)

		if result.Error != nil {
			continue
		}
	}

	// Insert Suppliers

	for i := 0; i < 5; i++ {
		supplier := model.Supplier{
			Name:        fmt.Sprintf("Supplier %d", i),
			PhoneNumber: fmt.Sprintf("0812%d", i),
			Address:     fmt.Sprintf("Alamat Supplier %d", i),
			PICName:     fmt.Sprintf("PIC Supplier %d", i),
		}

		result := pgDB.Create(&supplier)

		if result.Error != nil {
			continue
		}
		createdSupplierIDS = append(createdSupplierIDS, supplier.UUID)
	}

	if len(createdSupplierIDS) < 5 {
		panic("Supplier id is not enough")
	}

	// Insert Items
	for i := 0; i < 5; i++ {
		item := model.Item{
			Name:       fmt.Sprintf("Item %d", i),
			Category:   fmt.Sprintf("Category %d", i),
			Price:      11100.5,
			Unit:       "kg",
			SupplierID: createdSupplierIDS[i],
		}
		result := pgDB.Create(&item)

		if result.Error != nil {
			continue
		}

		createdItem = append(createdItem, item)
	}

	// Insert Products
	productCollection := mongodDb.Database("inventory_service").Collection("products")

	for i := 0; i < 5; i++ {
		ingredients := []model.Ingredient{}

		// Each product has 2 random ingredients
		for j := 0; j < 2; j++ {
			ingredients = append(ingredients, model.Ingredient{
				ItemID:   createdItem[j].UUID,
				ItemName: createdItem[j].Name,
				Quantity: (j + 1) * 10, // Example quantity
				Unit:     createdItem[j].Unit,
			})
		}

		product := model.Product{
			Name:        fmt.Sprintf("Product %d", i),
			Ingredients: ingredients,
		}

		_, err := productCollection.InsertOne(context.TODO(), product)
		if err != nil {
			fmt.Println("Error inserting product:", err)
			continue
		}
	}
}
