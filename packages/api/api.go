package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Event struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Datetime time.Time `json:"date"`
	Price    string    `json:"price"`
	Artists  []Artist  `json:"artists"`
	About    string    `json:"about"`
	Banners  []string  `json:"banners"`
}

type Artist struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
}

func CreateEvent(w http.ResponseWriter, req *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Problem in reading data !")
		return
	}
	event := Event{}
	err = json.Unmarshal([]byte(reqBody), &event)
	if err != nil {
		fmt.Fprintf(w, "Data not in proper format !")
		return
	}
	_, err = collection.InsertOne(context.TODO(), event)
	if err != nil {
		fmt.Fprintf(w, "Problem inserting to DB !")
		return
	}
	fmt.Fprintf(w, "Event created successfully")
	w.Close()

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
	id := req.URL.Query().Get("id")
	var result Event
	filter := bson.M{{"_id", id}}
	err = db.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Fprintf(w, "No such events found !")
	} else {
		jsonResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsonResult))
	}
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

	filter := bson.D{{"name", event}}

	update := bson.D{{"$set", bson.D{{"price", "50"}, {"city", "London"}}}}

	updateResult, err := db.Collection.UpdateOne(context.TODO(), filter, update)
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

	filter := bson.D{{"name", event}}
	deleteResult, err := db.Collection.DeleteOne(context.TODO(), filter)
	if deleteResult.DeletedCount == 0 {
		fmt.Fprintf(w, "No such events found !")
	} else {
		fmt.Fprintf(w, "Event successfully deleted !")
	}
}

func HandleEvent(w http.ResponseWriter, req *http.Request) {

	//id := req.URL.Query().Get("id")
	switch req.Method {
	//Read event data
	case http.MethodGet:
		getEvent(w, req)
	//Update event data
	case http.MethodPut:
		updateEvent(w, req)
	//Delete event data
	case http.MethodDelete:
		deleteEvent(w, req)
	default:
	}
}
