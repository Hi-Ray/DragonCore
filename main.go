package main

import (
	"dragonback/controllers/api/content/docs"
	"dragonback/controllers/api/home"
	"dragonback/lib/constants/repositories"
	"dragonback/lib/modules"
	"github.com/labstack/echo/v4"
)

func main() {
	modules.GitModule().EnsureRepo(repo.CDragonDocs)

	e := echo.New()
	home.Controller.Register(e)
	docs.Controller.Register(e)

	e.Logger.Fatal(e.Start(":3000"))
}