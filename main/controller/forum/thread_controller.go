package threadController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/forum_services"
)

func GetThread(c echo.Context) (err error) {
	result := forum_services.GetThreadPage()
	return c.String(http.StatusOK, result)
}

func GetThreadDetail(c echo.Context) (err error){
	res := forum_services.GetThreadDetail()
	return c.String(http.StatusOK, res)
}