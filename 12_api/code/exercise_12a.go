package main

import (
	"fmt"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

// create a route for /people
// write a function that uses fmt.Fprint that indicates the request was successful

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You've reached the people endpoint!")
}

func main() {
	fmt.Println(BaseURL)
	http.HandleFunc("/people", getPeople)
	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
