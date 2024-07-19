package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jain-vatsal/go-jwt/constants"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = GetDBInstance()

func GetDBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading the .env file", err)
	}

	mongoDbURI := os.Getenv(constants.MONGODB_URL)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbURI))
	if err != nil {
		log.Fatal("Error in setting up mongo client", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), constants.CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error in connecting to mongo client using context", err)
	}

	log.Printf("Connected to mongoDb\n")
	return client
}

func OpenCollection(client *mongo.Client, connectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(connectionName)
}
