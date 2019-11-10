package threadController

import (
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/forum_services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetThread(c echo.Context) (err error) {
	page := new(request.ThreadRequest)
	if err = c.Bind(page); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}
	result := forum_services.GetThreadPage(page)
	return c.String(http.StatusOK, result)
}

func GetThreadCategory(c echo.Context) (err error) {
	cat := new(request.ThreadCategoryRequest)
	if err = c.Bind(cat); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.GetThreadCategoryPage(cat)
	return c.String(http.StatusOK, result)
}

func GetThreadMaxPage(c echo.Context) (err error) {
	param := new(request.ThreadMaxPageRequest)
	if err = c.Bind(param); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}
	result := forum_services.GetMaxPage(param)
	return c.String(http.StatusOK, strconv.Itoa(result))
}

func CreateThread(c echo.Context) (err error) {
	r := new(request.CreateThreadRequest)

	usrname := c.Param("username")

	fmt.Println("user: ",r)

	r.MakerUsername = usrname

	fmt.Println("request: ", r)

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.CreateThread(r)

	return c.String(http.StatusOK, result)
}
