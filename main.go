package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitea.izolight.xyz/gabor/algodat/datastructures"
)

var stack = datastructures.NewStack(20)

func main() {
	stack.Push(1)
	stack.Push(9)
	stack.Push(-3)
	stack.Push(5)
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/stack", ViewStack)
	r.HandleFunc("queue", ViewQueue)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	ds := []string{"stack", "queue"}
	for _, d := range ds {
		fmt.Fprintf(w, "<a href='/%s'>%s</a><br>", d, d)
	}
}

func ViewStack(w http.ResponseWriter, r *http.Request) {
	element, err := stack.Top()
	if err != nil {
		return
	}
	s := stack.Size -1
	fmt.Fprint(w, "<table><caption>Stack</caption><tr><th>Position</th><th>Value</th><tr>")
	for {
		fmt.Fprintf(w, "<tr><td>%d</td><td>%d</td></tr>", s, element.Value)
		element, err = element.Next()
		if err != nil {
			break
		}
		s--
	}
	fmt.Fprint(w, "</table>")
}

func ViewQueue(w http.ResponseWriter, r *http.Request) {

}
