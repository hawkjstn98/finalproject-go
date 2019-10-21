package main

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
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


