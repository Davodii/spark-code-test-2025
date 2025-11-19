package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Storage system is in-memory so
// create a struct to hold to-do items
type ToDoItem struct {
	// Weird how to export fields they must be capitalised
	// Also, struct tags are needed to map the lowercase
	// JSON keys to the struct fields

	Title       string `json:"title"`
	Description string `json:"description"`
}

var toDolist []ToDoItem

func main() {
	// Adding some dummy date
	// toDolist = append(toDolist, ToDoItem{
	// 	Title:       "A",
	// 	Description: "B",
	// })

	// Set the http handler function
	http.HandleFunc("/", ToDoListHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Testing so just print
	// io.WriteString(w, "BELLO!\n")

	switch r.Method {
	case http.MethodOptions:
		// The application/json Content-Type header means
		// the browser does a preflight. Need to respond
		// with OK to allow for communication.

		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		// Encode the items array to JSON
		json.NewEncoder(w).Encode(toDolist)
	case http.MethodPost:
		// Parse the JSON into a ToDoItem struct
		var newItem ToDoItem

		dec := json.NewDecoder(r.Body)

		// For testing purposes, print that we have a POST request
		// fmt.Println("POSTing")

		err := dec.Decode(&newItem)
		if err != nil {
			// Retturn a 400 error if the JSON is invalid
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Store the new item in toDoList
		toDolist = append(toDolist, newItem)

		// Return the created item as JSON
		json.NewEncoder(w).Encode(newItem)
	default:
		// Method was not GET or POST
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
