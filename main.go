package main

import (
	"fmt"
	"log"
	"net/http"

	"gitea.izolight.xyz/gabor/algodat/datastructures"
	"github.com/gorilla/mux"
)

var stack = datastructures.NewStack(20)

func main() {
	stack.Push(1)
	stack.Push(9)
	stack.Push(-3)
	stack.Push(5)
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/stack", stack.View).Methods("GET")
	r.HandleFunc("/stack/push", stack.Add).Methods("POST")
	r.HandleFunc("/stack/pop", stack.Remove).Methods("POST")
	r.HandleFunc("queue", ViewQueue)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	ds := []string{"stack", "queue"}
	for _, d := range ds {
		fmt.Fprintf(w, "<a href='/%s'>%s</a><br>", d, d)
	}
}

func ViewQueue(w http.ResponseWriter, r *http.Request) {

}
