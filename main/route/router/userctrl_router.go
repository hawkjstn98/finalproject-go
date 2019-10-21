package router

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/controller"
	"github.com/labstack/echo"
)

func UCtrl(e *echo.Echo){

	e.GET("/getData", controller.GetAllData())
}
