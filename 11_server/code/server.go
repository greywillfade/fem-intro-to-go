package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Todo is a task that we need to do
type Todo struct {
	Title   string
	Content string
}

//PageVariables are variables sent to the HTML template
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome home!")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
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

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing eror: ", err)
	}

	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)
	fmt.Println("Server is running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) //note that Fatal is only called if an error occurs
}
