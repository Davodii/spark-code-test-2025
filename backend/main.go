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

	// Unique identifier for the to-do item
	ID int `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
}

var toDolist map[int]ToDoItem = make(map[int]ToDoItem)
var currentID int = 1

func main() {
	// Set the http handler function
	http.HandleFunc("/", ToDoListHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Testing so just print
	// io.WriteString(w, "BELLO!\n")

	switch r.Method {
	case http.MethodGet:
		// Convert the toDoList map to a slice
		items := make([]ToDoItem, 0, len(toDolist))
		for _, item := range toDolist {
			items = append(items, item)
		}

		// Encode the items array to JSON
		json.NewEncoder(w).Encode(items)
	case http.MethodPost:
		// Parse the JSON into a ToDoItem struct
		var newItem ToDoItem

		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&newItem)
		if err != nil {
			// Retturn a 400 error if the JSON is invalid
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Assign a unique ID to the new item
		newItem.ID = currentID
		currentID++

		// Store the new item in the toDoList map
		toDolist[newItem.ID] = newItem

		// Return the created item as JSON
		json.NewEncoder(w).Encode(newItem)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
