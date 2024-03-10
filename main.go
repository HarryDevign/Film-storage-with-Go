package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title string
	Director string
}

func main( ) { 
	fmt.Println("Hello World")

	h1 := func (w http.ResponseWriter, r *http.Request)  {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film {
			"Films": {
				{Title: "The Godfather", Director:  "Francis Ford Coppola"},
				{Title: "2001: A Space Odyssey", Director: "Stanley Kubrick"},
				{Title: "The Thing", Director:  "Curtis Hanson"},
			},
		}
		tmpl.Execute(w,  films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}