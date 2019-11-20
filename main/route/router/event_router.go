package router

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/controller/event"
	"github.com/labstack/echo/v4"
)

func EventRouter(e *echo.Echo) {
	e.GET("/getEventHome", event.Home)
	e.PUT("/createEvent/:username", event.CreateEvent)
	e.GET("/getEventDetail", event.DetailEvent)
	e.GET("/getMyEvent/:username/:latitude/:longitude", event.MyEvent)
	e.GET("/searchEvent/:username/:searchKey/:page/:latitude/:longitude", event.SearchEvent)
}