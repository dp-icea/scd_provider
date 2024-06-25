package store

import (
	"context"
	"icea_uss/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() *mongo.Client {
	conf := config.GetGlobalConfig()
	clientOptions := options.Client().ApplyURI("mongodb://" + conf.MongoUser + ":" + conf.MongoPassword + "@" + conf.MongoUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
