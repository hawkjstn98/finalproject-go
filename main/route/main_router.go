package route

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/route/router"
	"github.com/labstack/echo"
	"net/http"
)

func MainRouter(e *echo.Echo){
	e.GET("/", hello)

	router.UCtrl(e)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}