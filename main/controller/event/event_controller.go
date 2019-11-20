package event

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/request_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/services/event_services"
	"github.com/hawkjstn98/FinalProjectEnv/main/utility"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func MyEvent(c echo.Context) (err error) {
	username := c.Param("username")
	latitude := c.Param("latitude")
	longitude := c.Param("longitude")

	req := new(request.MyEventRequest)
	req.Username = username
	req.Latitude = latitude
	req.Longitude = longitude

	if err = c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	if req.Username == "" || req.Latitude == "" || req.Longitude == "" {
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	res := event_services.MyEventService(req)

	return c.String(http.StatusOK, res)
}

func SearchEvent(c echo.Context) (err error) {
	r := new(request.SearchEventRequest)

	usrname   := c.Param("username")
	search    := c.Param("searchKey")
	page 	  := c.Param("page")
	latitude  := c.Param("latitude")
	longitude := c.Param("longitude")

	r.SearchKey = search
	r.Username  = usrname
	r.Page,err = strconv.Atoi(page)
	r.Latitude = latitude
	r.Longitude = longitude

	if err != nil {
		if err = c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, request_constant.BadRequestError)
		}
		return c.String(http.StatusBadRequest, request_constant.BadRequestError)
	}

	result := event_services.SearchEvent(r)

	return c.String(http.StatusOK, result)
}
