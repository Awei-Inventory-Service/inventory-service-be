package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() (*mongo.Client, error) {

	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")

	if mongoUser == "" || mongoPassword == "" || mongoHost == "" || mongoPort == "" {
		log.Fatalf("Username / password / host / port can't be empty")
		return nil, errors.New("Check your Mongo env")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPassword, mongoHost, mongoPort)
	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
