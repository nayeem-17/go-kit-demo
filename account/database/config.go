package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const mongoSource = "mongodb://root:example@localhost:27017/"
var mongoSource = os.Getenv("MONGO_URI")

func GetDatabase(ctx context.Context, database_name string) *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(mongoSource).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB: ", err)
		os.Exit(-1)

	}
	fmt.Println("Connected to MongoDB!")
	return client.Database(database_name)
}
