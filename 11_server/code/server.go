package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Server is running on port :8080")
	// log.Fatal is a back-up wrapper for the inner function call in case http.ListenAndServe encounters errors
	// http.ListenAndServe is run regardless
	log.Fatal(http.ListenAndServe(":8080", nil))
}
