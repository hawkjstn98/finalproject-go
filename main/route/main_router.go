package route

import (
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/route/router"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"net/http"
	"os"
)

func MainRouter(e *echo.Echo) {
	e.GET("/", hello)
	e.File("/docs/swagger", "main/docs/docs.json")
	e.Static("/docs", "main/docs")
	router.UserRouter(e)
	router.ForumRouter(e)
	router.EventRouter(e)
	router.BookmarkRouter(e)

	path, exists := os.LookupEnv("PORT")

	if exists {
		// Print the value of the environment variable
		fmt.Print(path)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":1323"
	}

	e.Logger.Fatal(e.Start(port))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
