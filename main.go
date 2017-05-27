package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/stack", ViewStack)
	r.HandleFunc("queue", queue)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	datastructures := []string{"stack", "queue"}
	for _, d := range datastructures {
		fmt.Fprintf(w, "<a href='/%s'>%s</a><br>", d, d)
	}
}

func ViewStack(w http.ResponseWriter, r *http.Request) {
	stack := new(Stack)
}
