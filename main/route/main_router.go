package route

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/route/router"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"net/http"
)

// @title Swagger Final Project API
// @version 1.0
// @description .
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /v2
func MainRouter(e *echo.Echo) {
	e.GET("/", hello)
	e.File("/docs/swagger", "main/docs/docs.json")
	e.Static("/docs", "main/docs")
	router.UCtrl(e)
	router.FCtrl(e)

	e.Logger.Fatal(e.Start(":1323"))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
