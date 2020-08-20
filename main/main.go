package main

import (
	"log"
	"net/http"

	_ "github.com/Pankaj123Bhatt/Event-Management/packages/api"
	"github.com/Pankaj123Bhatt/Event-Management/packages/db"
)

func main() {
	port := ":8080"
	collection, err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	db.Collection = collection
	http.HandleFunc("/createEvent", api.createEvent)
	http.HandleFunc("/handleEvent", api.handleEvent)
	log.fatal(http.ListenAndServe(port, nil))
}
