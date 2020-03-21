package home

import (
	"dragonback/lib/models/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

var Controller = controller.New("").

	// Hello World
	Handler(controller.GET, "", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
