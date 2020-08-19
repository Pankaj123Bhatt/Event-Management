package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	collection, err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	db.Collection = collection

	/*NewServeMux()
	mux.HandleFunc("/events", apiH.ListEvents)
	mux.HandleFunc("/event/create", apiH.CreateEvent)
	mux.HandleFunc("/event", apiH.HandleEvent)*/

	http.HandleFunc("/createEvent", api.createEvent)
	http.HandleFunc("/handleEvent", api.handleEvent)
	http.ListenAndServe(port, nil)
}
