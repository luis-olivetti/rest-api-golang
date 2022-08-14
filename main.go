package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var people []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func populatePeople() {

	db, err := sql.Open("sqlite3", "./godb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT Id, FirstName, LastName FROM People")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName, lastName string

		err = rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{ID: strconv.Itoa(id), Firstname: firstName, Lastname: lastName})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	//people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
}

func main() {
	populatePeople()

	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")

	log.Println("Server running...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
