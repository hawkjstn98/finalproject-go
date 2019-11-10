package utility

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func GetHeader(c echo.Context, req string) (params map[string]*string, queries []string) {
	m := make(map[string]*string)
	queries = strings.Split(req, "|")
	for _, query := range queries {
		param := c.QueryParam(query)
		m[query] = &param
	}
	return m, queries
}

func Map(params map[string]*string, req []string, reqType interface{}) interface{} {
	if _, ok := reqType.(request.ThreadDetailRequest); ok {
		page, _ := strconv.Atoi(*params[req[1]])
		req := request.ThreadDetailRequest{
			ThreadID: *params[req[0]],
			Page:     page,
		}
		return req
	}
	if _, ok := reqType.(request.ThreadRequest); ok {
		page, _ := strconv.Atoi(*params[req[0]])
		req := request.ThreadRequest{
			Page: page,
		}
		return req
	}
	if _, ok := reqType.(request.ThreadCategoryRequest); ok {
		page, _ := strconv.Atoi(*params[req[1]])
		req := request.ThreadCategoryRequest{
			Category: *params[req[0]],
			Page: page,
		}
		return req
	}
	return nil
}