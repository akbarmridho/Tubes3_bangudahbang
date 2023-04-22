package routes

import (
	"backend/controllers"
	"github.com/labstack/echo/v4"
)

func QueryRoute(e *echo.Echo) {
	e.POST("/query", controllers.GetQueryHandler)
}
