package main

import (
	"log"

	"github.com/inventory-service/config"
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

	// initialize mongo db
	mongoDB, err := config.InitMongoDB()
	if err != nil {
		log.Fatalf("Error initializing mongo db, err=%v", err)
	}

	// for raw sql query
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("Error getting database connection pool, err=%v", err)
	// 	return
	// }

	// initialize redis

	router := InitRoutes(pgDB, mongoDB)

	router.Run(":8080")
}