package database

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	_ = godotenv.Load(".env")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
}

func OpenCollection(collectionName string) (*mongo.Collection, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	return client.Database("cluster0").Collection(collectionName), nil
}
