package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Todo is a todo with a title and content
type Todo struct {
	Title   string
	Content string
}

// PageVariables are variables sent to the html template
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	// Fprint prints to the external writer, can be shown on the web page in browser instead of the console
	// fmt.Fprint(w, "Todos!")

	pageVariables := PageVariables{
		PageTitle: "Get Todos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, pageVariables)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	fmt.Println("Server is running on port :8080")
	// log.Fatal is a back-up wrapper for the inner function call in case http.ListenAndServe encounters errors
	// http.ListenAndServe is run regardless
	log.Fatal(http.ListenAndServe(":8080", nil))
}
