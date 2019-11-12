package router

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/controller/bookmark"
	"github.com/labstack/echo/v4"
)

func BookmarkRouter(e *echo.Echo) {
	e.PUT("/bookmark/add", bookmark.Bookmark)
	e.PUT("/bookmark/remove", bookmark.RemoveBookmark)
}