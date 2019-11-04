package router

import (
	forumController "github.com/hawkjstn98/FinalProjectEnv/main/controller/forum"
)

func FCtrl(e *echo.Echo) {
	e.GET("/getThreadPage", forumController.GetThread)
}
