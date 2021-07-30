package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json: "ID"`
	Title       string `json: "Title"`
	Date        string `json: "Date"`
	Description string `json: "Description`
	Category    string `json: "Category"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Bleachers Concert",
		Date:        "10/12/21",
		Description: "Bleachers at The Union",
		Category:    "Concert",
	},
	{
		ID:          "2",
		Title:       "CHVRCHES Concert",
		Date:        "12/10/21",
		Description: "CHVRCHES at The Union",
		Category:    "Concert",
	},
	{
		ID:          "3",
		Title:       "Take The Sadness Out Of Saturday Night",
		Date:        "7/30/21",
		Description: "Bleachers Album Release",
		Category:    "Music",
	},
	{
		ID:          "4",
		Title:       "Screen Violence",
		Date:        "8/27/21",
		Description: "CHVRCHES Album Release",
		Category:    "Music",
	},
	{
		ID:          "5",
		Title:       "Deathloop",
		Date:        "9/14/21",
		Description: "Deathloop Release",
		Category:    "PS5",
	},
	{
		ID:          "6",
		Title:       "Metroid Dread",
		Date:        "10/8/21",
		Description: "Metroid Dread Release",
		Category:    "Nintendo",
	},
	{
		ID:          "7",
		Title:       "Switch OLED",
		Date:        "10/8/21",
		Description: "Switch OLED Release",
		Category:    "Nintendo",
	},
	{
		ID:          "8",
		Title:       "Waitress with Sara Bareilles",
		Date:        "10/2/21",
		Description: "Waitress at the Barrymore Theatre",
		Category:    "Concert",
	},
}
var newId = 9

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	newEvent.ID = fmt.Sprint(newId)
	newId += 1
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Date = updatedEvent.Date
			singleEvent.Description = updatedEvent.Description
			singleEvent.Category = updatedEvent.Category
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
