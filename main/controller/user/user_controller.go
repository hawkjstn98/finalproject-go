package userController

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) (err error) {
	u := new(user.User)

	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	//return c.String(http.StatusOK, response)
	return nil
}
