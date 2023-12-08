package main

import (
	"mini-project/database"
	"mini-project/routes"
	"mini-project/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	util.InitConfig()

	database.InitDB()

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} '${uri}' [${status}] (${latency_human})\n",
	}))

	routes.RoutesInit(e)

	e.Logger.Fatal(e.Start(":" + util.Cfg.PORT))
}
