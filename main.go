package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkEnv(key string) string {
	data := os.Getenv(key)
	if data == "" {
		log.Fatalf("Error: Environment variable %s not set\n", key)
	}
	return data
}

func envDefault(key, defaultValue string) string {
	data := os.Getenv(key)
	if data == "" {
		return defaultValue
	}
	return data
}

func connectToMongoDB(mongoURL string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return client
}

func main() {
	// Get environment variables of "MONGO_DB_NAME" and "MONGO_URL"
	// If they are not set, use default values
	mongoDBName := checkEnv("MONGO_DB_NAME")
	mongoURL := checkEnv("MONGO_URL")
	mongoCollection := envDefault("MONGO_COLLECTION", "delete_me")

	if mongoCollection == "delte_me" {
		log.Println("Warning: Using default collection name 'delete_me'")
	}

	// Connect to MongoDB
	client := connectToMongoDB(mongoURL)
	defer client.Disconnect(context.Background())

	// Create MongoDB Database
	err := client.Database(mongoDBName).CreateCollection(context.Background(), mongoCollection, nil)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println("Database", "'"+mongoDBName+"'", "with collection", "'"+mongoCollection+"'", "created successfully")
}
