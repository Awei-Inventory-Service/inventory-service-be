package seeding

import (
	"fmt"
	"os"

	constant "github.com/inventory-service/lib/constants"
	"gorm.io/gorm"
)

func MainSeed(pgDB *gorm.DB) {
	fmt.Println(os.Getenv(constant.FeatureFlagSeedInitialData))
	if os.Getenv(constant.FeatureFlagSeedInitialData) == "true" {
		Seed(pgDB)
	}

	if os.Getenv(constant.FeatureFlagSeedStockTransactionData) == "true" {
		SeedStockTransaction(pgDB)
	}
}
