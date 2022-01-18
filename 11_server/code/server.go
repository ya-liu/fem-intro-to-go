package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

func todos(w http.ResponseWriter, r *http.Request) {
	// Fprint prints to the external writer, can be shown on the web page in browser instead of the console
	// fmt.Fprint(w, "Todos!")

	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", todos)
	fmt.Println("Server is running on port :8080")
	// log.Fatal is a back-up wrapper for the inner function call in case http.ListenAndServe encounters errors
	// http.ListenAndServe is run regardless
	log.Fatal(http.ListenAndServe(":8080", nil))
}
