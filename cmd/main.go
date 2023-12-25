package main

import (
	"go+html/endpoints"
	"log"
	"text/template"

	"github.com/labstack/echo"
)

func main() {
	_, err := template.ParseFiles(
		"./public/views/index.html",
		"./public/views/name.html",
	)

	if err != nil {
		log.Fatalf("could not initialize templates: %v", err)
	}

	e := echo.New()
	e.GET("/", endpoints.HandleIndex)
	e.GET("/name", endpoints.HandleName)

	e.Logger.Fatal(e.Start(":8000"))
}
