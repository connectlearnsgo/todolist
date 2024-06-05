package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type List struct {
	Name  string
	Items []string
}

var lists []List

func main() {
	lists = []List{
		List{
			Name: "Groceries",
			Items: []string{
				"Apples",
				"Oranges",
				"Bananas",
			},
		},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", showHome)
	mux.HandleFunc("/todos/", handleTodos)
	mux.HandleFunc("/todos", handleTodos)
	// mux.HandleFunc("/todos/", showTodoEdit)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/todos")
	if path != "/" {
		switch r.Method {
		case http.MethodGet:
			showTodo(w, r)
		case http.MethodPut:
			updateTodo(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
		return
	}

	switch r.Method {
	case http.MethodGet:
		showTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func showHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Everybody!")
}

func showTodos(w http.ResponseWriter, r *http.Request) {
	for _, list := range lists {
		fmt.Fprintf(w, "%s\n", list.Name)
	}
}

func showTodo(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, _ := strconv.Atoi(idStr)
	if id < len(lists) {
		fmt.Fprintf(w, "%v", lists[id])
	} else {
		http.Error(w, "List not found", http.StatusNotFound)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	list := List{
		Name: name,
	}

	lists = append(lists, list)

	fmt.Fprintf(w, "List created: %s\n", list.Name)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, _ := strconv.Atoi(idStr)
	if id < len(lists) {
		lists[id].Name = r.FormValue("name")
		http.Redirect(w, r, fmt.Sprintf("/todos/%d", id), http.StatusSeeOther)
		return
	}

	http.Error(w, "List not found", http.StatusNotFound)
}
