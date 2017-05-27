package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gitea.izolight.xyz/gabor/algodat/datastructures"
	"github.com/gorilla/mux"
)

var stack = datastructures.NewStack(20)
var templates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/viewstack.html"))

type viewData struct {
	Title string
	Data  interface{}
}
type stackData struct {
	Position int
	Value    int
}

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
		fmt.Fprintf(w, "Could not get top of stack: %v", err)
	}
	var values []stackData
	s := stack.Size - 1
	for {
		values = append(values, stackData{s, element.Value})
		element, err = element.Next()
		if err != nil {
			break
		}
		s--
	}
	data := viewData{"Stack", values}
	err = templates.ExecuteTemplate(w, "viewstack", data)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
}

func ViewQueue(w http.ResponseWriter, r *http.Request) {

}
