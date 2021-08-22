package main

import (
	"fmt"
	"log"
	"net/http"
)

//Gives a http response based on given http request
func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the RESTAPI!")
	fmt.Println("Endpoint Hit: homePage")
}

//Will match URL path (just localhost in this case) and provide a http response from homePage func
func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}
