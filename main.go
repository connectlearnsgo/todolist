package main

import (
	"fmt"
	"log"
	"net/http"
)

type List struct {
	Name  string
	Items []string
}

var lists []List

func main() {
	lists = []List{}

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", dispatch)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		showLists(w, r)
	case http.MethodPost:
		createList(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func showLists(w http.ResponseWriter, r *http.Request) {
	for _, list := range lists {
		fmt.Fprintf(w, "%s\n", list.Name)
	}
}

func createList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	list := List{
		Name: name,
	}

	lists = append(lists, list)

	fmt.Fprintf(w, "List created: %s\n", list.Name)
}
