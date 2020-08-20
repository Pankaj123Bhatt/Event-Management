package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/Pankaj123Bhatt/Event-Management/packages/db"
	"go.mongodb.org/mongo-driver/bson"
)

ttype Event struct {
	ID       string  `json:"id" bson:"id"`
	Title    string  `json: "title" bson:"title,omitempty"`
	Language string  `json: "language" bson:"language,omitempty"`
	Genre    string  `json:"genre" bson:"genre,omitempty"`
	Date     string  `json:"date" bson:"date"`
	About    string    `json:"about"`
	Time     string  `json: "time" bson:"time,omitempty"`
	Price    string  `json: "price" bson:"price,omitempty"`
	Artist   []Artist `json:"artist" bson:"artist,omitempty" `
	Banners  []string  `json:"banners,omitempty"`
}

type Artist struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Image string `json:"image" bson:"image,omitempty"`
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
	fmt.Fprintf(w, "Event created successfully !")
	

}

func getEvent(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Problem in reading data !")
		return
	}
	var event_id string
	err = json.Unmarshal([]byte(reqBody), &event_id)
	if err != nil {
		fmt.Fprintf(w, "Data not in proper format !")
		return
	}
	//id := event.ID
	var result Event
	filter := bson.D{{"_id", event_id}}
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
		fmt.Fprintf(w, "Problem in reading data !")
		return
	}
	var event Event
	err = json.Unmarshal([]byte(reqBody), &event)
	if err != nil {
		fmt.Fprintf(w, "Data not in proper format !")
		return
	}
	id := req.URL.Query().Get("id")
	filter := bson.D{{"_id", id}}
	

	update := bson.D{{"$set", event}}

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
		fmt.Fprintf(w, "Problem in reading data !")
		return
	}
	var event_id string
	err = json.Unmarshal([]byte(reqBody), &event_id)
	if err != nil {
		fmt.Fprintf(w, "Data not in proper format !")
		return
	}

	filter := bson.M{{"_id", event_id}}
	deleteResult, err := db.Collection.DeleteOne(context.TODO(), filter)
	if deleteResult.DeletedCount == 0 {
		fmt.Fprintf(w, "No such events found !")
	} else {
		fmt.Fprintf(w, "Event successfully deleted !")
	}
}

func HandleEvent(w http.ResponseWriter, req *http.Request) {

	
	switch req.Method {
	
	case http.MethodGet:
		getEvent(w, req)
	
	case http.MethodPut:
		updateEvent(w, req)
	
	case http.MethodDelete:
		deleteEvent(w, req)

	default:
		fmt.Fprintf(w, "No such Handle event API found !")
	}
}
