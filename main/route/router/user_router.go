package router

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/controller"
	userController "github.com/hawkjstn98/FinalProjectEnv/main/controller/user"
	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo) {
	e.GET("/getData", controller.GetAllData)
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
	e.PUT("/addUpdateGameList/:username", userController.AddOrUpdateGameList)
	e.PUT("/addUpdatePhoneNumber/:username", userController.AddOrUpdatePhoneNumber)
	e.GET("/getUserData/:username", userController.GetUserData)
	e.PUT("/addUpdateProfileImage/:username", userControlle)
}
