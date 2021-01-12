package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseURL = Config("DB_URL")
	databaseName = Config("DB_NAME")
	DB mongo.Database
)

func connectToMongo() {

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	fmt.Println("Connected to MongoDB.")
	DB = *client.Database(databaseName)
}

func init() {
	connectToMongo()
}
