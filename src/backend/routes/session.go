package routes

import (
	"backend/controllers"
	"github.com/labstack/echo/v4"
)

func SessionRoute(e *echo.Echo) {
	e.POST("/session", controllers.GetSessionHandler)
}
