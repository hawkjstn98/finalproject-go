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
	page := new(request.ThreadRequest)
	if err = c.Bind(page); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}
	result := forum_services.GetThreadPage(page)
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
	cat := new(request.ThreadCategoryRequest)
	if err = c.Bind(cat); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := forum_services.GetThreadCategoryPage(cat)
	return c.String(http.StatusOK, result)
}

func GetThreadMaxPage(c echo.Context) (err error) {
	param := new(request.ThreadCategoryRequest)
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
