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
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

type AllPeople struct {
	// this key called People doesn't exist on the json object
	// use a json flag to bridge the data name in json and what we call the data in go syntax
	People []Person `json:"results"`
}

// create a route for /people
// write a function that uses fmt.Fprint that indicates the request was successful

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "You've reached the people endpoint!")
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to get Star Wars people")
	}

	// fmt.Println(res)

	// ioutil is used to parse the http response for json data
	// the bytes variable was originally named data
	// after printing to the console, the returned value is an array of bytes like [5 99 108 101]
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse http response")
	}

	// fmt.Println(bytes)
	// fmt.Println(string(bytes))
	// after printing the stringifed bytes, the actual data is stored in results

	var people AllPeople

	// works like json.parse to turn the bytes data into a variable of choice, the 2nd argument

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, person := range people.People {
		person.getHomeworld()
		fmt.Println(person)
	}
}

func main() {
	// fmt.Println(BaseURL)
	http.HandleFunc("/people", getPeople)
	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
