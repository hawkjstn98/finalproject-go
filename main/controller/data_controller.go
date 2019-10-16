package controller

import (
	"encoding/json"
	"finalproject/main/repository"
	"github.com/labstack/echo"
	"net/http"
)

func GetAllData(c echo.Context) error {
	user := repository.LoadAllUserData()
	response, _ := json.Marshal(user)
	return c.String(http.StatusOK, string(response))
}
