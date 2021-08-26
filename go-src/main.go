package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

//The structure for the factsz
type Fact struct {
	ID       string `json:"ID"`
	FactType string `json:"factType"`
	Content  string `json:"content"`
}

var Facts []Fact
var db *sql.DB
var err error

//When calling to GET all facts we will encode it in json then return all fatcs to the response
//GET all request
func returnAllFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Facts []Fact

	result, err := db.Query("SELECT * FROM Facts")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var fact Fact
		err := result.Scan(&fact.ID, &fact.FactType, &fact.Content)
		if err != nil {
			panic(err.Error())
		}
		Facts = append(Facts, fact)
	}
  
	fmt.Println("Endpoint Hit: returnAllFacts")
	json.NewEncoder(w).Encode(Facts)
}

//Gives specific fact based off ID given in URL
//GET specific fact request
func returnSingleFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, fact := range Facts {
		if fact.ID == key {
			json.NewEncoder(w).Encode(fact)
		}
	}
}

//Will be used to let the user create a new fact for the DB
//POST request
func createNewFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var fact Fact
	_ = json.NewDecoder(r.Body).Decode(&fact)
	fact.ID = "7" //Note we generate a UUID for the fact
	Facts = append(Facts, fact)
	json.NewEncoder(w).Encode(fact)
}

//Will be used to let the user update one of the facts in the DB
//PUT request
func updateFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for index, fact := range Facts {
		fmt.Printf("%v", key)
		if fact.ID == key {
			Facts = append(Facts[:index], Facts[index+1:]...)
			var fact Fact
			_ = json.NewDecoder(r.Body).Decode(&fact)
			fact.ID = key //Keep same ID
			Facts = append(Facts, fact)
			json.NewEncoder(w).Encode(fact)
			break
		}
	}
}

//Will be used to let the user delete a fact from the DB
//DELETE request
func deleteFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for index, fact := range Facts {
		if fact.ID == key {
			Facts = append(Facts[:index], Facts[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(Facts)
}

//Gives a http response based on given http request
func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the RESTAPI!")
	fmt.Println("Endpoint Hit: homePage")
}

//Will match URL path (just localhost in this case) and provide a http response from homePage func
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllFacts).Methods("GET")
	myRouter.HandleFunc("/fact/{id}", returnSingleFact).Methods("GET")
	myRouter.HandleFunc("/fact", createNewFact).Methods("POST")
	myRouter.HandleFunc("/fact/{id}", updateFact).Methods("PUT")
	myRouter.HandleFunc("/fact/{id}", deleteFact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	//Connect to the mysql db
	db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/test_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	Facts = []Fact{
		{ID: "1", FactType: "Random", Content: "This is random fact"},
		{ID: "2", FactType: "Random", Content: "This is random fact2"},
	}
	handleRequest()
}
