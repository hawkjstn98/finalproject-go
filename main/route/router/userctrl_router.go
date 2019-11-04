package router

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/controller"
	userController "github.com/hawkjstn98/FinalProjectEnv/main/controller/user"
	forumController "github.com/hawkjstn98/FinalProjectEnv/main/controller/forum"
	"github.com/labstack/echo/v4"
)

func UCtrl(e *echo.Echo) {
	e.GET("/getData", controller.GetAllData)
	e.POST("/register", userController.Register)
	e.GET("/getThreadPage", forumController.GetThread)
}
