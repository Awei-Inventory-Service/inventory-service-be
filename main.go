package main

import (
	"log"

	"github.com/inventory-service/config"
	"github.com/inventory-service/seeding"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	// initialize pg db
	pgDB, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database, err=%v", err)
		return
	}

	// for raw sql query
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("Error getting database connection pool, err=%v", err)
	// 	return
	// }

	// initialize redis

	router := InitRoutes(pgDB)

	// UNCOMMENT INI KLO GAK MAU SEEDING
	seeding.Seed(pgDB)
	router.Run(":8080")
}
