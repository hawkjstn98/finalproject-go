package main

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	route.MainRouter(e)


	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}


