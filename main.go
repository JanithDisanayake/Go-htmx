package main

import (
	"go+htmx/models"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/films", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/films.html"))
		films := map[string][]models.Film{
			"Films": {
				{Title: "The Goadfather", Director: "Francis Ford Copppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	})

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/header.html"))
		tmpl.Execute(w, nil)
	})


	log.Fatal(http.ListenAndServe(":8000", nil))
}
