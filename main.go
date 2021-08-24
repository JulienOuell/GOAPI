package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//The structure for the factsz
type Fact struct {
	ID       string `json:"ID"`
	FactType string `json:"factType"`
	Content  string `json:"content"`
}

var Facts []Fact

//When calling to GET all facts we will encode it in json then return all fatcs to the response
func returnAllFacts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllFacts")
	json.NewEncoder(w).Encode(Facts)
}

func returnSingleFact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)
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
	myRouter.HandleFunc("/all", returnAllFacts)
	myRouter.HandleFunc("/fact/{id}", returnSingleFact)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Facts = []Fact{
		{ID: "1", FactType: "Random", Content: "This is random fact"},
		{ID: "2", FactType: "Random", Content: "This is random fact2"},
	}
	handleRequest()
}
