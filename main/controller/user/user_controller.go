package userController

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/user_services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) (err error) {
	u := new(request.RegisterRequest)

	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.RegisterUser(u)

	return c.String(http.StatusOK, result)
}
