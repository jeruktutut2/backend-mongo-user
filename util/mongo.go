package util

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jeruktutut2/backend-mongo-user/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(configurationMongo configuration.Mongo) (mongoDatabase *mongo.Database, client *mongo.Client) {
	credentials := options.Credential{
		AuthSource: configurationMongo.Database,
		Username:   configurationMongo.Username,
		Password:   configurationMongo.Password,
	}
	optionsClient := options.Client()
	optionsClient.ApplyURI("mongodb://" + configurationMongo.Host + ":" + strconv.Itoa(configurationMongo.Port)).SetAuth(credentials)
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
	return mongoDatabase, client
}

func CloseMongoDbConnection(client *mongo.Client) {

	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Error close mongodb connection: %v", err)
	}
	fmt.Println("Connection mongodb closed")
}
