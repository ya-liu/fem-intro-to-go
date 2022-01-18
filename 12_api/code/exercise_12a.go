package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

type Planet struct {
	Name       string
	Terrain    string
	Population string
}

type Person struct {
	Name         string
	HomeworldURL string
	Homeworld    Planet
}

type AllPeople struct {
	People []Person
}

// create a route for /people
// write a function that uses fmt.Fprint that indicates the request was successful

func getPeople(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "You've reached the people endpoint!")
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to get Star Wars people")
	}

	fmt.Println(res)

	// ioutil is used to parse the http response for json data
	// the bytes variable was originally named data
	// after printing to the console, the returned value is an array of bytes like [5 99 108 101]
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse http response")
	}

	fmt.Println(bytes)

	var people AllPeople

	// works like json.parse to turn the bytes data into a variable of choice, the 2nd argument
	json.Unmarshal(bytes, &people)
}

func main() {
	// fmt.Println(BaseURL)
	http.HandleFunc("/people", getPeople)
	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
