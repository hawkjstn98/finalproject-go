package event

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/event_services"
	"github.com/hawkjstn98/FinalProjectEnv/main/utility"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Home(c echo.Context) error {
	m, queries := utility.GetHeader(c, request_constant.EventHomeRequest)
	mappedReq := utility.Map(m, queries, request.EventHomeRequest{})
	req, ok := mappedReq.(request.EventHomeRequest)
	if !ok {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res, err := event_services.GetEventHome(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}
	return c.String(http.StatusOK, res)
}

func CreateEvent(c echo.Context) (err error) {
	usrname := c.Param("username")

	r := new(request.CreateEventRequest)

	if err = c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	var eventInsert event.EventInsert

	eventInsert.Timestamp = time.Now()
	eventInsert.Name = r.Name
	eventInsert.MakerUsername = usrname
	eventInsert.Type = r.Type
	eventInsert.Games = r.Games
	eventInsert.Category = r.Category
	eventInsert.Description = r.Description
	eventInsert.Site = r.Site
	eventInsert.DateStart = r.DateStart
	eventInsert.DateEnd = r.DateEnd
	eventInsert.Latitude = r.Latitude
	eventInsert.Longitude = r.Longitude
	eventInsert.Poster = r.Poster

	result := event_services.CreateEvent(eventInsert)

	return c.String(http.StatusOK, result)
}

func DetailEvent(c echo.Context) (err error) {
	m, queries := utility.GetHeader(c, request_constant.EventDetailRequest)
	mappedReq := utility.Map(m, queries, request.EventDetailRequest{})
	req, ok := mappedReq.(request.EventDetailRequest)
	if !ok {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res, err := event_services.EventDetail(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, request_constant.InternalServerError+" "+err.Error())
	}
	return c.String(http.StatusOK, res)
}