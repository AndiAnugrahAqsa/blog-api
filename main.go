package main

import (
	"mini-project/database"
	"mini-project/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.RoutesInit(e)

	e.Logger.Fatal(e.Start(":2020"))
}
