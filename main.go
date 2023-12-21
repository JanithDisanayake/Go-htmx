package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film {
			"Films": {
				{Title: "The Goadfather", Director: "Francis Ford Copppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
