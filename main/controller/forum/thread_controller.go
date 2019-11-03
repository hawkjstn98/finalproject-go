package threadController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/forum_services"
	"log"
)

func GetThread(c echo.Context) (err error) {
	log.Println("Masuk disini")
	result := forum_services.GetThreadPage()
	return c.String(http.StatusOK, result)
}