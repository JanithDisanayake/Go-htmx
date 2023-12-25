package endpoints

import (
	"go+html/models"

	"github.com/labstack/echo"
)

func HandleIndex(c echo.Context) error {
	films := []models.Film{
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "Blade Runner", Director: "Ridley Scott"},
		{Title: "The Thing", Director: "John Carpenter"},
	}

	return c.Render(200, "index.html", films)
}

func HandleName(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}
