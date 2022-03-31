package util

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection() (mongoDatabase *mongo.Database) {
	credentials := options.Credential{
		AuthSource: "backendusermongo",
		Username:   "admin",
		Password:   "12345",
	}
	optionsClient := options.Client()
	optionsClient.ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	client, err := mongo.NewClient(optionsClient)
	if err != nil {
		log.Fatal("new client connection mongo err:", err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("connection mongo err:", err)
	}

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	mongoDatabase = client.Database("backendusermongo")
	return mongoDatabase
}
