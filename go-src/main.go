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

var db *sql.DB
var err error

//When calling to GET all facts we will encode it in json then return all fatcs to the response
//GET all request
func returnAllFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var facts []Fact

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
		facts = append(facts, fact)
	}
  
	fmt.Println("Endpoint Hit: returnAllFacts")
	json.NewEncoder(w).Encode(facts)
}

//Gives specific fact based off ID given in URL
//GET specific fact request
func returnSingleFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	result, err := db.Query("SELECT * FROM Facts WHERE id = ?", key)
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()

	var fact Fact
	for result.Next() {
		err = result.Scan(&fact.ID, &fact.FactType, &fact.Content)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(fact)
}

//Will be used to let the user create a new fact for the DB
//POST request
func createNewFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	  
	var fact Fact
	var stmt *sql.Stmt
	json.NewDecoder(r.Body).Decode(&fact)
	if fact.ID == "" {
		stmt, err = db.Prepare("INSERT INTO Facts(FactType, Content) VALUES(?, ?)")
	} else {
		stmt, err = db.Prepare("INSERT INTO Facts(ID, FactType, Content) VALUES(?, ?, ?)")
	}
	if err != nil {
    	panic(err.Error())
  	}

	if fact.ID == "" {
		_, err = stmt.Exec(fact.FactType, fact.Content)
	} else {
		_, err = stmt.Exec(fact.ID, fact.FactType, fact.Content)
	}
	if err != nil {
    	panic(err.Error())
  	}
	
	w.WriteHeader(http.StatusCreated) // 201 to client
	json.NewEncoder(w).Encode(fact)
}

//Will be used to let the user update one of the facts in the DB
//PUT request
func updateFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	var fact Fact
	var stmt *sql.Stmt
	json.NewDecoder(r.Body).Decode(&fact)

	stmt, err = db.Prepare("UPDATE Facts SET FactType = ?, Content = ? WHERE id = ?")
	if err != nil {
    	panic(err.Error())
  	}
	
	fact.ID = key
	_, err = stmt.Exec(fact.FactType, fact.Content, fact.ID)
	if err != nil {
    	panic(err.Error())
  	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(fact)
}

//Will be used to let the user delete a fact from the DB
//DELETE request
func deleteFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	stmt, err := db.Prepare("DELETE FROM Facts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(key)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusNoContent) //204 to the client
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
	handleRequest()
}
