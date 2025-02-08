package main

import (
	"flag"
	"log"

	"github.com/inventory-service/config"
	"github.com/inventory-service/seeding"
	"github.com/joho/godotenv"
)

func main() {
	seedFlag := flag.Bool("seed", false, "Run database seeding")
	flag.Parse()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	// Initialize PostgreSQL database
	pgDB, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database, err=%v", err)
		return
	}

	// Check if the seeding flag is passed
	if *seedFlag {
		log.Println("Running database seeding...")
		seeding.SeedStockTransaction(pgDB)
		log.Println("Seeding completed!")
		return
	}

	// Initialize and start the router
	router := InitRoutes(pgDB)
	router.Run(":8080")
}