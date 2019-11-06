package threadController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/forum_services"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"strconv"
)

func GetThread(c echo.Context) (err error) {
	result := forum_services.GetThreadPage()
	return c.String(http.StatusOK, result)
}

func GetThreadDetail(c echo.Context) (err error){
	req := new(request.ThreadDetailRequest)
	if err = c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res, err := forum_services.GetThreadDetail(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError)
	}
	return c.String(http.StatusOK, res)
}

func GetThreadCategory(c echo.Context) (err error) {
	cat := new(request.ThreadCategoryRequest)
	if err = c.Bind(cat); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.GetThreadCategoryPage(cat)
	return c.String(http.StatusOK, result)
}

func GetThreadMaxPage(c echo.Context) (err error){
	param := new(request.ThreadMaxPageRequest)
	if err = c.Bind(param); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}
	result := forum_services.GetMaxPage(param)
	return c.String(http.StatusOK, strconv.Itoa(result))
}