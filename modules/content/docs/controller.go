package docs

import (
"dragonback/lib/models/controller"
"dragonback/lib/models/repo"
"github.com/labstack/echo/v4"
"net/http"
)

var handler = controller.New("/content/docs").

	// returns an index of items
	Handler(controller.GET, "", func () echo.HandlerFunc {
		files, _ := repo.New(repo.CDragonDocs).Documentation()

		return func(c echo.Context) error {
			var paths []string
			for _, file := range files {
				paths = append(paths, file.URL())
			}
			return c.JSON(http.StatusOK, paths)
		}
	}()).

	// returns a markdown page
	Handler(controller.GET, "/:item", func () echo.HandlerFunc {
		files, _ := repo.New(repo.CDragonDocs).Documentation()

		return func(c echo.Context) error {
			for _, file := range files {
				if file.URL() == c.Param("item") {
					c.Response().Header().Set(echo.HeaderContentType, "text/md")
					return c.String(http.StatusOK, file.Content())
				}
			}
			return c.NoContent(404)
		}
	}())
