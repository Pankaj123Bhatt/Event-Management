package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	Name  string
	Price string
	City  string
}

func getEvent(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var event string
	err = json.Unmarshal([]byte(reqBody), &event)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://user1234:1234@cluster0.fun2x.mongodb.net/EventManager?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("EventManager").Collection("Event")

	var result Event
	filter := bson.D{{"name", event}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Fprintf(w, "No such events found !")
	} else {
		jsonResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsonResult))
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	event := Event{}
	err = json.Unmarshal([]byte(reqBody), &event)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://user1234:1234@cluster0.fun2x.mongodb.net/EventManager?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("EventManager").Collection("Event")

	_, err = collection.InsertOne(context.TODO(), event)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Event created successfully")
}

func updateEvent(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var event string
	err = json.Unmarshal([]byte(reqBody), &event)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://user1234:1234@cluster0.fun2x.mongodb.net/EventManager?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("EventManager").Collection("Event")

	filter := bson.D{{"name", event}}

	update := bson.D{{"$set", bson.D{{"price", "50"}, {"city", "London"}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if updateResult.ModifiedCount == 0 {
		fmt.Fprintf(w, "No such events found !")
	} else {
		fmt.Fprintf(w, "Event successfully updated !")
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var event string
	err = json.Unmarshal([]byte(reqBody), &event)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://user1234:1234@cluster0.fun2x.mongodb.net/EventManager?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("EventManager").Collection("Event")

	filter := bson.D{{"name", event}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if deleteResult.DeletedCount == 0 {
		fmt.Fprintf(w, "No such events found !")
	} else {
		fmt.Fprintf(w, "Event successfully deleted !")
	}
}

func handleRequests() {

	http.HandleFunc("/getEvent", getEvent)
	http.HandleFunc("/createEvent", createEvent)
	http.HandleFunc("/updateEvent", updateEvent)
	http.HandleFunc("/deleteEvent", deleteEvent)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequests()
}
