package main

import (
	"mini-project/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	e.Logger.Fatal(e.Start(":2020"))
}
