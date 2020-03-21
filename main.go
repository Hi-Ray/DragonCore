package main

import (
	"dragonback/controllers/content/docs"
	"dragonback/controllers/home"
	"dragonback/lib/models/repo"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	err := repo.New(repo.CDragonDocs).Clone()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	home.Controller.Register(e)
	docs.Controller.Register(e)

	e.Logger.Fatal(e.Start(":3000"))
}