package main

import (
	"dragonback/modules"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	modules.Bundler.Register(e)
	e.Logger.Fatal(e.Start(":3000"))
}
