package controller

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/user_services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllData(c echo.Context) error {
	response := user_services.GetAllUserData()
	return c.String(http.StatusOK, response)
}
