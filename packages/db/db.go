package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection mongo.Collection

func Init() (mongo.Collection, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := "user1234"    //read from file
	pass := "1234"        //read from file
	cluster := "cluster0" //read from file
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+user+":"+pass+"@"+cluster+".fun2x.mongodb.net/EventManager?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}
	collection := client.Database("EventManager").Collection("Event")
	return collection, nil
}
