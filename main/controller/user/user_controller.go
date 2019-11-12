package userController

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/user_services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) (err error) {
	r := new(request.RegisterRequest)

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.RegisterUser(r)

	return c.String(http.StatusOK, result)
}

func Login(c echo.Context) (err error) {
	r := new(request.LoginRequest)

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.LoginUser(r)

	return c.String(http.StatusOK, result)
}

func AddOrUpdateGameList(c echo.Context) (err error) {
	r := new(request.AddOrUpdateGameListRequest)

	usrname := c.Param("username")

	r.Username = usrname

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.AddOrUpdateGameList(r)

	return c.String(http.StatusOK, result)
}

func AddOrUpdatePhoneNumber(c echo.Context) (err error) {
	r := new(request.AddOrUpdatePhoneRequest)

	usrname := c.Param("username")

	r.Username = usrname

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.AddOrUpdatePhone(r)

	return c.String(http.StatusOK, result)
}

func GetUserData(c echo.Context) (err error) {

	usrname := c.Param("username")

	if "" == usrname {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.GetUserData(usrname)

	return c.String(http.StatusOK, result)
}

func AddUpdateProfileImage(c echo.Context) (err error) {
	request := new(request.AddOrUpdateProfileImage)

	usrname := c.Param("username")

	request.Username = usrname;

	if "" == usrname {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := user_services.AddOrUpdateProfileImage()

	return c.String(http.StatusOK, result)
}
