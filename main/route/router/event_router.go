package router

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/controller/event"
	"github.com/labstack/echo/v4"
)

func EventRouter(e *echo.Echo) {
	e.GET("/getEventHome", event.Home)
	e.PUT("/createEvent/:username", event.CreateEvent)
	e.GET("/getEventDetail", event.DetailEvent)
}