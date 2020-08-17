package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Event struct {
	Name  string
	Price string
	City  string
}

func main() {

	var operationType int
	fmt.Println("Welcome to Event Management Application !")
	for {

		fmt.Println("Please press :\n 1 - Creating event.\n 2 - Getting information.\n 3 - Updating event.\n 4 - Delete event.\n 5 - Exit the application.\n")
		fmt.Printf("Enter your choice : ")
		fmt.Scanf("%d", &operationType)

		switch operationType {

		case 1: // Create Event

			var eventName, eventPrice, eventCity string
			fmt.Printf("Please enter the event name : ")
			fmt.Scanf("%s", &eventName)
			fmt.Printf("Please enter the event price : ")
			fmt.Scanf("%s", &eventPrice)
			fmt.Printf("Please enter the event venue : ")
			fmt.Scanf("%s", &eventCity)

			jsonData := Event{eventName, eventPrice, eventCity}
			jsonValue, _ := json.Marshal(jsonData)
			request, _ := http.NewRequest("POST", "http://localhost:8080/createEvent", bytes.NewBuffer(jsonValue))
			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
			}

		case 2: // Get Event

			var jsonData string
			fmt.Printf("Please enter the event name : ")
			fmt.Scanf("%s", &jsonData)
			jsonValue, _ := json.Marshal(jsonData)
			request, _ := http.NewRequest("GET", "http://localhost:8080/getEvent", bytes.NewBuffer(jsonValue))
			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
			}

		case 3: // Update Event

			var jsonData string
			fmt.Printf("Please enter the event name : ")
			fmt.Scanf("%s", &jsonData)
			jsonValue, _ := json.Marshal(jsonData)
			request, _ := http.NewRequest("POST", "http://localhost:8080/updateEvent", bytes.NewBuffer(jsonValue))
			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
			}

		case 4: // Delete Event

			var jsonData string
			fmt.Printf("Please enter the event name : ")
			fmt.Scanf("%s", &jsonData)
			jsonValue, _ := json.Marshal(jsonData)
			request, _ := http.NewRequest("GET", "http://localhost:8080/deleteEvent", bytes.NewBuffer(jsonValue))
			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
			}

		case 5: // Quit application
			fmt.Println("Exiting Application...")
			os.Exit(1)
		default:
			fmt.Println("Please enter valid choice !")

		}

	}

}
