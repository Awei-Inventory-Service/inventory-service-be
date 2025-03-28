package main

import (
	"log"

	"github.com/inventory-service/config"
	"github.com/inventory-service/seeding"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		log.Fatalf("Error loading environment files")
		return
	}

	// Initialize PostgreSQL database
	pgDB, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database, err=%v", err)
		return
	}

	// Seed data
	seeding.MainSeed(pgDB)

	// Initialize and start the router
	router := InitRoutes(pgDB)

	router.Run("0.0.0.0:8080")
}
