package main

import (
	"fmt"
	"log"
	"net/http"

	"gitea.izolight.xyz/gabor/algodat/datastructures"
	"github.com/gorilla/mux"
)

var stack = datastructures.NewStack(10)
var queue = datastructures.NewQueue(10)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/stack", stack.View).Methods("GET")
	r.HandleFunc("/stack/push", stack.Push).Methods("POST")
	r.HandleFunc("/stack/pop", stack.Pop).Methods("POST")
	r.HandleFunc("/queue", queue.View).Methods("GET")
	r.HandleFunc("/queue/enqueue", queue.Enqueue).Methods("POST")
	r.HandleFunc("/queue/dequeue", queue.Dequeue).Methods("POST")
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
