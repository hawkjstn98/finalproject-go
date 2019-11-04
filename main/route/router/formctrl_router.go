package router

import (
	forumController "github.com/hawkjstn98/FinalProjectEnv/main/controller/forum"
	"github.com/labstack/echo/v4"
)

func FCtrl(e *echo.Echo) {
	e.GET("/getThreadPage", forumController.GetThread)
}
