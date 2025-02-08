package seeding

import (
	"os"

	constant "github.com/inventory-service/lib/constants"
	"gorm.io/gorm"
)

func MainSeed(pgDB *gorm.DB) {
	if os.Getenv(constant.FeatureFlagSeedInitialData) == "true" {
		Seed(pgDB)
	}

	if os.Getenv(constant.FeatureFlagSeedStockTransactionData) == "true" {
		SeedStockTransaction(pgDB)
	}
}
