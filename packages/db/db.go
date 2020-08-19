package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*type mongoCollection struct {
	collection mongo.Collection
}*/
var Collection mongo.Collection

func Init() (mongo.Collection, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://user1234:1234@cluster0.fun2x.mongodb.net/EventManager?retryWrites=true&w=majority", //read from file
	))
	if err != nil {
		return nil, err
	}
	collection := client.Database("EventManager").Collection("Event")
	return collection, nil
}
