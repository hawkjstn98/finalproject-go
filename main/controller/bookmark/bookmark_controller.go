package bookmark

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/event_services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Bookmark(c echo.Context) error {
	r := new(request.BookmarkRequest)

	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res, err := event_services.Bookmark(r)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}
	return c.String(http.StatusOK, res)
}

func RemoveBookmark(c echo.Context) error {
	r := new(request.BookmarkRequest)

	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res, err := event_services.RemoveBookmark(r)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}
	return c.String(http.StatusOK, res)
}
