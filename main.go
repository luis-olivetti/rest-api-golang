package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var people []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func populatePeople() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
}

func main() {
	populatePeople()

	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")

	log.Println("Server running...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
