package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"scd_provider/scd/dss"
)

func CreateOir(oir dss.OperationalIntent) error {
	oirCollection := db().Database("scd").Collection("operational_intent")
	_, err := oirCollection.InsertOne(context.TODO(), oir)
	if err != nil {
		return err
	}
	return nil
}

func GetOir(id string) (dss.OperationalIntent, error) {
	oirCollection := db().Database("scd").Collection("operational_intent")
	var result dss.OperationalIntent

	err := oirCollection.FindOne(context.TODO(), bson.M{"reference.id": id}).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
