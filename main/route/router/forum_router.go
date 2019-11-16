package router

import (
	forumController "github.com/hawkjstn98/FinalProjectEnv/main/controller/forum"
	"github.com/labstack/echo/v4"
)

func ForumRouter(e *echo.Echo) {
	e.GET("/getThreadPage", forumController.GetThread)
	e.GET("/getThreadDetail", forumController.GetThreadDetail)
	e.GET("/getThreadCategoryPage", forumController.GetThreadCategory)
	e.PUT("/createThread/:username", forumController.CreateThread)
}
