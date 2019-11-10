package threadController

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/utility"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/forum_services"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
)

func GetThread(c echo.Context) (err error) {
	m, queries := utility.GetHeader(c, request_constant.ThreadRequest)
	mappedReq := utility.Map(m, queries, request.ThreadRequest{})
	req, ok := mappedReq.(request.ThreadRequest)
	if !ok {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.GetThreadPage(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}
	return c.String(http.StatusOK, result)
}

func GetThreadDetail(c echo.Context) (err error) {
	m, queries := utility.GetHeader(c, request_constant.ThreadDetailRequest)
	mappedReq := utility.Map(m, queries, request.ThreadDetailRequest{})
	req, ok := mappedReq.(request.ThreadDetailRequest)
	if !ok {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res, err := forum_services.GetThreadDetail(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}
	return c.String(http.StatusOK, res)
}

func GetThreadCategory(c echo.Context) (err error) {
	m, queries := utility.GetHeader(c, request_constant.ThreadCategoryRequest)
	mappedReq := utility.Map(m, queries, request.ThreadCategoryRequest{})
	req, ok := mappedReq.(request.ThreadCategoryRequest)
	if !ok {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.GetThreadCategoryPage(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}

	return c.String(http.StatusOK, result)
}

func GetThreadMaxPage(c echo.Context) (err error) {
	m, queries := utility.GetHeader(c, request_constant.ThreadCategoryRequest)
	mappedReq := utility.Map(m, queries, request.ThreadCategoryRequest{})
	req, ok := mappedReq.(request.ThreadCategoryRequest)
	if !ok {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.GetMaxPage(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}

	return c.String(http.StatusOK, result)
}

func CreateThread(c echo.Context) (err error) {
	r := new(request.CreateThreadRequest)

	usrname := c.Param("username")

	r.MakerUsername = usrname

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.CreateThread(r)

	return c.String(http.StatusOK, result)
}
