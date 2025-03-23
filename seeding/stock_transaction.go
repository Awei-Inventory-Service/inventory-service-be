package seeding

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inventory-service/app/model"
	"gorm.io/gorm"
)

func SeedStockTransaction(pgDB *gorm.DB) {
	var item model.Item
	var branches []model.Branch
	var users []model.User

	// Fetch Item ID 1
	if err := pgDB.Where("name = ?", "Item 1").First(&item).Error; err != nil {
		fmt.Println("Error fetching item:", err)
		return
	}

	// Fetch branches (ensuring at least 2 exist)
	if err := pgDB.Find(&branches).Error; err != nil || len(branches) < 2 {
		fmt.Println("Error fetching branches or not enough branches")
		return
	}

	// Fetch users (for issuer)
	if err := pgDB.Find(&users).Error; err != nil || len(users) == 0 {
		fmt.Println("Error fetching users")
		return
	}

	// Seed random stock transactions while ensuring positive stock balance
	for i := 0; i < 5; i++ {
		// Generate a random OUT quantity (between 5 and 20)
		outQuantity := rand.Intn(16) + 5 // Ensures it's at least 5

		// Ensure IN quantity is always more than OUT to keep stock positive
		inQuantity := outQuantity + rand.Intn(10) + 5 // Adds buffer to prevent negative stock

		// OUT Transaction (Stock Moving Out from Branch 1 to Branch 2)
		outTransaction := model.StockTransaction{
			BranchOriginID:      branches[0].UUID, // First branch as origin
			BranchDestinationID: branches[1].UUID, // Second branch as destination
			ItemID:              item.UUID,        // Item 1
			IssuerID:            users[0].UUID,    // First user as issuer
			Type:                "OUT",
			Quantity:            outQuantity,
			Cost:                float64(outQuantity) * 1000.0, // Example cost calculation
			Reference:           fmt.Sprintf("OUT-REF-%d", i),
			Remarks:             "Stock moved out",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		}

		// IN Transaction (Stock Moving Back from Branch 2 to Branch 1)
		inTransaction := model.StockTransaction{
			BranchOriginID:      branches[1].UUID, // Second branch as origin
			BranchDestinationID: branches[0].UUID, // First branch as destination
			ItemID:              item.UUID,        // Item 1
			IssuerID:            users[0].UUID,    // First user as issuer
			Type:                "IN",
			Quantity:            inQuantity,                   // Always greater than outQuantity
			Cost:                float64(inQuantity) * 1000.0, // Adjusted cost
			Reference:           fmt.Sprintf("IN-REF-%d", i),
			Remarks:             "Stock moved in",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		}

		// Insert OUT Transaction
		if err := pgDB.Create(&outTransaction).Error; err != nil {
			fmt.Println("Error inserting OUT stock transaction:", err)
			continue
		}

		// Insert IN Transaction
		if err := pgDB.Create(&inTransaction).Error; err != nil {
			fmt.Println("Error inserting IN stock transaction:", err)
			continue
		}
	}

	fmt.Println("Stock transactions (IN & OUT) seeded successfully!")
}
